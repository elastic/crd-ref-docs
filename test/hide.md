# API Reference

## Packages
- [webapp.test.k8s.elastic.co/common](#webapptestk8selasticcocommon)
- [webapp.test.k8s.elastic.co/v1](#webapptestk8selasticcov1)


## webapp.test.k8s.elastic.co/common

Package common contains common API Schema definitions



#### CommonString

_Underlying type:_ _string_





_Appears in:_
- [GuestbookSpec](#guestbookspec)
- [GuestbookStatus](#guestbookstatus)




## webapp.test.k8s.elastic.co/v1

Package v1 contains API Schema definitions for the webapp v1 API group

### Resource Types
- [Embedded](#embedded)
- [Guestbook](#guestbook)
- [GuestbookList](#guestbooklist)
- [Underlying](#underlying)



#### Embedded









| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `webapp.test.k8s.elastic.co/v1` | | |
| `kind` _string_ | `Embedded` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `a` _string_ |  |  |  |
| `x` _string_ |  |  |  |


#### Embedded1







_Appears in:_
- [Embedded](#embedded)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `x` _string_ |  |  |  |


#### EmbeddedX







_Appears in:_
- [Embedded](#embedded)
- [Embedded1](#embedded1)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `x` _string_ |  |  |  |


#### Guestbook



Guestbook is the Schema for the guestbooks API.



_Appears in:_
- [GuestbookList](#guestbooklist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `webapp.test.k8s.elastic.co/v1` | | |
| `kind` _string_ | `Guestbook` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[GuestbookSpec](#guestbookspec)_ |  | \{ page:1 \} |  |


#### GuestbookEntry



GuestbookEntry defines an entry in a guest book.



_Appears in:_
- [GuestbookSpec](#guestbookspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `name` _string_ | Name of the guest (pipe \| should be escaped) |  | MaxLength: 80 <br />Pattern: `0*[a-z0-9]*[a-z]*[0-9]` <br /> |
| `time` _[Time](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#time-v1-meta)_ | Time of entry |  |  |
| `comment` _string_ | Comment by guest. This can be a multi-line comment.<br />Like this one.<br />Now let's test a list:<br />* a<br />* b<br /><br />Another isolated comment.<br /><br />Looks good? |  | Pattern: `0*[a-z0-9]*[a-z]*[0-9]*` <br /> |
| `rating` _[Rating](#rating)_ | Rating provided by the guest |  | Maximum: 5 <br />Minimum: 1 <br /> |


#### GuestbookHeader

_Underlying type:_ _string_

GuestbookHeaders are strings to include at the top of a page.



_Appears in:_
- [GuestbookSpec](#guestbookspec)



#### GuestbookList



GuestbookList contains a list of Guestbook.





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `webapp.test.k8s.elastic.co/v1` | | |
| `kind` _string_ | `GuestbookList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Guestbook](#guestbook) array_ |  |  |  |


#### GuestbookSpec



GuestbookSpec defines the desired state of Guestbook.



_Appears in:_
- [Guestbook](#guestbook)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `page` _[PositiveInt](#positiveint)_ | Page indicates the page number | 1 | Minimum: 1 <br /> |
| `entries` _[GuestbookEntry](#guestbookentry) array_ | Entries contain guest book entries for the page |  |  |
| `selector` _[LabelSelector](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#labelselector-v1-meta)_ | Selector selects something |  |  |
| `headers` _[GuestbookHeader](#guestbookheader) array_ | Headers contains a list of header items to include in the page |  | MaxItems: 10 <br />UniqueItems: true <br /> |
| `certificateRef` _[SecretObjectReference](https://gateway-api.sigs.k8s.io/references/spec/#gateway.networking.k8s.io/v1beta1.SecretObjectReference)_ | CertificateRef is a reference to a secret containing a certificate |  |  |
| `str` _[CommonString](#commonstring)_ |  |  |  |
| `enum` _[MyEnum](#myenum)_ | Enumeration is an example of an aliased enumeration type |  | Enum: [MyFirstValue MySecondValue] <br /> |




#### MyEnum

_Underlying type:_ _string_



_Validation:_
- Enum: [MyFirstValue MySecondValue]

_Appears in:_
- [GuestbookSpec](#guestbookspec)

| Field | Description |
| `MyFirstValue` | MyFirstValue is an interesting value to use<br /> |
| `MySecondValue` | MySecondValue is what you use when you can't use MyFirstValue<br /> |


#### PositiveInt

_Underlying type:_ _integer_



_Validation:_
- Minimum: 1

_Appears in:_
- [GuestbookSpec](#guestbookspec)



#### Rating

_Underlying type:_ _integer_

Rating is the rating provided by a guest.

_Validation:_
- Maximum: 5
- Minimum: 1

_Appears in:_
- [GuestbookEntry](#guestbookentry)





#### Underlying



Underlying tests that Underlying1's underlying type is Underlying2 instead of string.





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `webapp.test.k8s.elastic.co/v1` | | |
| `kind` _string_ | `Underlying` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `a` _[Underlying1](#underlying1)_ |  | b | MaxLength: 10 <br /> |


#### Underlying1

_Underlying type:_ _[Underlying2](#underlying2)_

Underlying1 has an underlying type with an underlying type

_Validation:_
- MaxLength: 10

_Appears in:_
- [Underlying](#underlying)



#### Underlying2

_Underlying type:_ _string_

Underlying2 is a string alias

_Validation:_
- MaxLength: 10

_Appears in:_
- [Underlying1](#underlying1)



