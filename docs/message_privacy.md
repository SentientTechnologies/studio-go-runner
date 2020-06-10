# Message Encryption

This section describes the message encryption, and signing features of the runner.  Message payloads are described in the docs/interface.md file.  Encryption, and signing is only supported within Kubernetes deployments.  The reason for this is that standalone runners cannot be secured and have shared secrets without the isolation provided by Kubernetes.

Encrypted payloads use a hybrid cryptosystem, [please click for a detailed description](https://en.wikipedia.org/wiki/Hybrid_cryptosystem).

Message signing uses Ed25519 signing as defined by RFC8032, more information can be found at[https://ed25519.cr.yp.to/](https://ed25519.cr.yp.to/).

Ed25519 certificate SHA1 fingerprints, not intended to be cryptographicaly secure, will be used by clients to assert identity, confirmed by successful verification.a  Verification still relies on a full public key.

<!--ts-->

Table of Contents
=================

* [Message Encryption](#message-encryption)
* [Table of Contents](#table-of-contents)
* [Introduction](#introduction)
* [Encryption](#encryption)
  * [Key creation by the cluster owner](#key-creation-by-the-cluster-owner)
* [Mount secrets into runner deployment](#mount-secrets-into-runner-deployment)
  * [Message format](#message-format)
* [Signing](#signing)
  * [Signing deployment](#signing-deployment)
    * [Manual insertion](#manual-insertion)
    * [Automatted insertion](#automatted-insertion)
* [Python StudioML configuration](#python-studioml-configuration)
<!--te-->

# Introduction

This document describes encryption of Request messages sent by StudioML clients to the runner.

Encryption of messages has two tiers, the first tier is a Public-key scheme that has the runner employ a private key and a public key that is given to experimenters using the python or other client software.

The concerns to users of the system is to obtain from the computer cluster owner the public key, and only the public key.  The public key can then be made accessible to the client for securing the messages exchanged with the runner compute instances.

The compute cluster owner will be resposible for generating the public-private key pair and manging the integrity of the private key.  They will also be responsible for distribution of the public key to any experiments, or users of the system.

The client encrypts a per message secret that is encrypted using the public key, and prepended to a payload that contains the request message encrypted using the secret.

# Encryption

## Key creation by the cluster owner

The owner of the compute cluster is responsible for the generation of key pair for use with the message encryption.  The following commands show the creation of the key pairs.

```
echo -n "PassPhrase" > secret_phrase
ssh-keygen -t rsa -b 4096 -f studioml_message -C "Message Encryption Key" -N "PassPhrase"
ssh-keygen -f studioml_message.pub -e -m PEM > studioml_message.pub.pem
cp studioml_message studioml_message.pem
ssh-keygen -f studioml_message.pem -e -m PEM -p -P "PassPhrase" -N "PassPhrase"
```

The private key file and the passphrase should be considered as valuable secrets for your organization that MUST be protected and cared for appropriately.

Once the keypair has been created they can be loaded into the Kubernetes runner cluster using the following commands:

```
kubectl create secret generic studioml-runner-key-secret --from-file=ssh-privatekey=studioml_message.pem --from-file=ssh-publickey=studioml_message.pub.pem
kubectl create secret generic studioml-runner-passphrase-secret --from-file=ssh-passphrase=secret_phrase
```

The passphrase is kept in a seperate secret to enable RBAC access to be used to isolate the two pieces of knowledge should your secrets management procedures call for this.

The public PEM key MUST be the only file delivered to client side users of StudioML in PEM Key file format, for example:

```
-----BEGIN RSA PUBLIC KEY-----
MIICCgKCAgEAtZurOEVuT9bhjiUWX7U8EFxL8oMGWSLXf4M6QBsJ5TljtSqyIxvI
kXiQDLIpJXY8KRmiR9RghGopvB5NfAMLZtfwozuju2NtnSn0UPI+6O4ED6TfDP5F
eta/6tUKAuvxVwF5Yvr7en1qnbv4L86vqeukrn/gIPTb7LlsFjt6uHlxA6xTAun/
HfRKlBiWR5rIi/fwuUMmTGpAcCa8s5Gqfla28FfsknGOipy4Vw4Mt7f93ke1dHN+
dY/J2TpCm/GNJuFaHc4EgHE8uw+jU6uBgpZAJSIzK5dxYniEjZS93CWxs2HN8dmV
wEqleT02agWW4cfa13X3Lz1YoQkCjYtSqB8Y2KjT1q7sSll0HExWV58kFPk9FmIy
JniMLcLFzAxGDM5UgtmsdSYmqN49vlqOejxfYxy6GrKXrkRGCDuQKyb2m/WQLXGU
8cGqwuVpN/JNWjiG4+NaxWRzfE2Yk4gbhcYqXRocNMlidG0Sx/xrFTFln86lmGJ1
RCse6jv3beENf5lfrz4ddAzAssjTivmlZgJCTK2oROT3WPI/G6CaBQadt13XkQLW
hAZDbnsZMhOVH3/UiQJ6DwgV0yK5FND4jkbHM3GWGNLRIrnL9F0I8c1p9X2oCx6T
plgCug3iz5cE9+G2455Y1vaVMBEKSm1REhsdTYzPBV/yXPpPR4lUCmkCAwEAAQ==
-----END RSA PUBLIC KEY-----
```

A single key pair is used to encrypt all requests on the cluster at this time.  A future feature is envisioned to allow multiple key pairs.

When the runner is run the secrets are mounted into the container that Kubernetes is managing.  This is done using the deployment yaml.  When performing deployments the yaml should be reviewed for runner pod, and their runner container to ensure that the secrets are available and that they are mounted.  If these secrets are not loaded into the cluster the runner pod should remain in a pending state.

# Mount secrets into runner deployment

Secrets used by the runner will be mounted into the runner pod using the Kubernetes deployment pod resource definition.  An example of this is provided within the sample AWS CPU runner that can be found in the [../examples/aws/cpu/deployment.yaml](../examples/aws/cpu/deployment.yaml) file.

Two mounts will be created firstly for the keyfiles, secondly for the passphrase.  These two are split to allow for RBAC to be employed in the cluster should you want it.  The motivation is that you might want to divide ownership between two parties for the private key and the and avoid revealing one of these to the other.

If you wish to use encrypted traffic exclusively be sure to remove the ```CLEAR_TEXT_MESSAGES: "true"``` entry from your ConfigMap entries in the yaml.

In any event the yaml need to mount these secrets appears as follows:

```
apiVersion: apps/v1
kind: Deployment
metadata:
 name: studioml-go-runner-deployment
 labels:
   app: studioml-go-runner
spec:
 ...
 template:
   ...
   spec:
      ...
      containers:
      - name: studioml-go-runner
        ...
        volumeMounts:
        - name: message-encryption
          mountPath: "/runner/certs/message/encryption"
          readOnly: true
        - name: encryption-passphrase
          mountPath: "/runner/certs/message/passphrase"
          readOnly: true
        ...
      volumes:
        ...
        - name: message-encryption
          secret:
            optional: false
            secretName: studioml-runner-key-secret
            items:
            - key: ssh-privatekey
              path: ssh-privatekey
            - key: ssh-publickey
              path: ssh-publickey
        - name: encryption-passphrase
          secret:
            optional: false
            secretName: studioml-runner-passphrase-secret
            items:
            - key: ssh-passphrase
              path: ssh-passphrase
```

## Message format

The encrypted\_data block contains two comma seperated Base64 strings.  The first string contains a symmetric key that is encrypted using RSA-OAEP with a key length of 4096 bits, and the sha256 hashing algorithm. The second field contains the JSON string for the Request message that is first encrypted using a NaCL SecretBox encryption and then encoded as Base64.

The encryption works in two steps, first the secretbox based symmetric shared key is generated for every message by the source generating the message.  The data within the messages is encrypted with the symmetric key.  The symmetric key is then encrypted and placed at the front of the message using an asymmetric key.  This has the following effects:

The sender can decrypt the payload if they retain their original symmetric key.
The sender can not decrypt the symmetric key, once it is placed encrypted into the payload
The legitimate runner if able to access the RSA PEM private key can decrypt the asymmetric key, and only then can subsequently decrypt the Request in the payload.
Evesdropping software cannot decrypt the asymmetricly encrypted secretbox key and so cannot decrypt the rest of the payload.

# Signing

Message signing is a way of protecting the runner receiving messages from processing spoofed requests.  To prevent this the runner can be configured to read public key information from Kubernetes secrets and then to use this to validate messages that are being received.  The configuration information for the runner signing keys is detailed in the next section.

Signing is only supported in Kubernetes deployments.

The portion of the message that is signed is the Base64 representation of the encrypted payload.

Message signing uses Ed25519 signing as defined by RFC8032, more information can be found at[https://ed25519.cr.yp.to/](https://ed25519.cr.yp.to/).

Ed25519 certificate SHA256 fingerprints, not intended to be cryptographicaly secure, will be used by clients to assert identity, confirmed by successful verification. Verification of messages sent to the runner relies on a public key supplied by the experimenter.  The follow example shows how an experimenter would go about creating a private public key pair suitable for signing:

```
ssh-keygen -t ed25519 -f studioml_signing -P ""
ssh-keygen -l -E sha256 -f studioml_signing.pub
256 SHA256:BB+StMfwvv/8Dutb0i1QpdBL171Fg/Fd3ODebi+NX74 kmutch@awsdev (ED25519)
```

The finger print can be extracted and sent to the cluster administrator, from the last line of the above output.

Having generated a key pair the PUBLIC key file should be transmitted to the administrators of any runner compute clusters that will be used.  Along with sending the key the experimenter should decide in conjunction with their community the queue name prefixes they will be assigned to use exclusively. The queue name prefixes should be passed to the administrators with the public key pem file.

Queue name prefixes should be a minimum of four characters to include the queue technology being used with the underscore, for example 'rmq_', or 'sqs_' to use the public key on all four queues.

If you send the request via email you might compose something like the following to send:

```
Hi,

I would like to add/replace a signing verification key for any queues on the 54.123.10.5 Rabbit MQ Server for our cluster with the prefix of 'rmq_cpu_andrei_'.

They public key I wish to use is:

ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIFITo06Pk8sqCMoMHPaQiQ7BY3pjf7OE8BDcsnYozmIG kmutch@awsdev

Our fingerprint is:

SHA256:BB+StMfwvv/8Dutb0i1QpdBL171Fg/Fd3ODebi+NX74

Thanks,
Andrei
```

The above should provide enough information to the administrator to apply your key to the system and reply using email confirming the key has been added.

Once a message signing public key has been assigned any messages on related queue MUST have a valid signature attached to messages otherwise they will be rejected.

## Signing deployment

Before starting any addition of message signing keys the cluster administrator must check that the request being sent originated from a pre-nominated sender.

Signing keys can be injected into the compute cluster using Kubernetes secrets.  The runners in a cluster will use a secret in the same namespace called 'studioml-signing' for extracting signing keys.  The addition of new keys is via the addition of data items within the secrets resource via the kubectl apply command. Changes or additions to signing keys are propogated via the mounted resource within the runner pods, see [Mounted Secrets are updated automatically](https://kubernetes.io/docs/concepts/configuration/secret/#mounted-secrets-are-updated-automatically).

Using the example, above, then a secret data item can be added to the studio signing secrets using a command such as the following example workflow shows:

```
$ export KUBECTL_CONFIG=~/.kube/my_cluster.config
$ export KUBECTLCONFIG=~/.kube/my_cluster.config
$ kubectl get secrets
NAME                                TYPE                                  DATA   AGE
default-token-qps8p                 kubernetes.io/service-account-token   3      11s
docker-registry-config              Opaque                                1      11s
release-github-token                Opaque                                1      11s
studioml-runner-key-secret          Opaque                                2      11s
studioml-runner-passphrase-secret   Opaque                                1      11s
studioml-signing                    Opaque                                1      11s
```
```
$ kubectl get secrets studioml-signing -o=yaml
apiVersion: v1
data:
  info: RHVtbXkgU2VjcmV0IHNvIHJlc291cmNlIHJlbWFpbnMgcHJlc2VudA==
kind: Secret
metadata:
  annotations:
    kubectl.kubernetes.io/last-applied-configuration: |
      {"apiVersion":"v1","data":{"info":"RHVtbXkgU2VjcmV0IHNvIHJlc291cmNlIHJlbWFpbnMgcHJlc2VudA=="},"kind":"Secret","metadata":{"annotations":{},"name":"studioml-signing","namespace":"default"},"type":"Opaque"}
  creationTimestamp: "2020-05-15T22:05:26Z"
  managedFields:
  - apiVersion: v1
    fieldsType: FieldsV1
    fieldsV1:
      f:data:
        .: {}
        f:info: {}
      f:metadata:
        f:annotations:
          .: {}
          f:kubectl.kubernetes.io/last-applied-configuration: {}
      f:type: {}
    manager: kubectl
    operation: Update
    time: "2020-05-15T22:05:26Z"
  name: studioml-signing
  resourceVersion: "790034"
  selfLink: /api/v1/namespaces/ci-go-runner-kmutch/secrets/studioml-signing
  uid: bc13f78d-199b-4afb-8b3a-31b6ea486c8e
type: Opaque
```

This next line will take the public key that was emailed to you and convert it into Base 64 format ready to be inserted into the Kubernetes secret input encoding.

```
$ item=`cat << EOF | base64 -w 0
ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIFITo06Pk8sqCMoMHPaQiQ7BY3pjf7OE8BDcsnYozmIG kmutch@awsdev
EOF
`
```

### Manual insertion

If you do not have the jq tool installed you will now have to manually edit the secret using the following command:

```
$ kubectl edit secrets studioml-signing
```

Now manually insert a yaml line after the info: item so that things appear as follows:

```
  1 # Please edit the object below. Lines beginning with a '#' will be ignored,
  2 # and an empty file will abort the edit. If an error occurs while saving this file will be
  3 # reopened with the relevant failures.
  4 #
  5 apiVersion: v1
  6 data:
  7   info: RHVtbXkgU2VjcmV0IHNvIHJlc291cmNlIHJlbWFpbnMgcHJlc2VudA==
  8   rmq_cpu_andrei_: c3NoLWVkMjU1MTkgQUFBQUMzTnphQzFsWkRJMU5URTVBQUFBSUZJVG8wNlBrOHNxQ01vTUhQYVFpUTdCWTNwamY3T0U4QkRjc25Zb3ptSUcga211dGNoQGF3c2Rldgo=
  9 kind: Secret
 10 metadata:
 11   annotations:
... [redacted] ...
```

Now use the ':wq' command to exit the editor and have the secret updated inside the cluster.

### Automatted insertion

Using the jq command the new secret can be inserted into the secret using the following:

```
kubectl get secret mysecret -o json | jq --arg item= "${item}" '.data["rmq_cpu_andrei_"]=$item' | kubectl apply -f -
```

# Python StudioML configuration

In order to use experiment payload encryption with the Python-based StudioML client,
the StudioML section of experiment configuration must specify
a path to the public key file in PEM format. If a path is not specified,
the experiment payload will be submitted unencrypted, in plain text form.

If a StudioML configuration is provided as part of the enclosing completion service configuration, in .hocon format, it would include the following (example):

```
{
   ...
   "studio_ml_config": {
         ...
         "public_key_path": "/home/user/keys/my-key.pub.pem",
         ...
   }
   ...
}
```

another possibility is:

```
{
   ...
   "studio_ml_config": {
         ...
         "public_key_path": ${PUBLIC_KEY_PATH},
         ...
   }
   ...
}
```

For the base StudioML configuration, in .yaml format, specifying the public key for encryption would look like:

```
public_key_path: /home/user/keys/my-key.pub.pem
```

If you wish to use message signing to prove that queue messages you send to the cluster are from a genuine sender then an additional option can be specified, for example:

```
{
   ...
   "studio_ml_config": {
         ...
         "public_key_path": "/home/user/keys/my-key.pub.pem",
         "signing_key_path": "/home/user/keys/studioml_signing",
         ...
   }
   ...
}
```

Copyright © 2019-2020 Cognizant Digital Business, Evolutionary AI. All rights reserved. Issued under the Apache 2.0 license.