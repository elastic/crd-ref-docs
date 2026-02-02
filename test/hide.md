# API Reference

Here is a template value: `v1`.

## Packages
- [webapp.test.k8s.elastic.co/common](#webapp-test-k8s-elastic-co-common)
- [webapp.test.k8s.elastic.co/v1](#webapp-test-k8s-elastic-co-v1)


## <a id="webapp-test-k8s-elastic-co-common">webapp.test.k8s.elastic.co/common</a>

Package common contains common API Schema definitions

*Important: This package is special and should be treated differently.*


#### <a id="github-com-elastic-crd-ref-docs-api-common-commonstring">CommonString</a>

_Underlying type:_ _string_





_Appears in:_
- [GuestbookSpec](#github-com-elastic-crd-ref-docs-api-v1-guestbookspec)
- [GuestbookStatus](#github-com-elastic-crd-ref-docs-api-v1-guestbookstatus)




## <a id="webapp-test-k8s-elastic-co-v1">webapp.test.k8s.elastic.co/v1</a>

Package v1 contains API Schema definitions for the webapp v1 API group

### Resource Types
- [Embedded](#github-com-elastic-crd-ref-docs-api-v1-embedded)
- [Guestbook](#github-com-elastic-crd-ref-docs-api-v1-guestbook)
- [GuestbookList](#github-com-elastic-crd-ref-docs-api-v1-guestbooklist)
- [Underlying](#github-com-elastic-crd-ref-docs-api-v1-underlying)



#### <a id="github-com-elastic-crd-ref-docs-api-v1-embedded">Embedded</a>









| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `webapp.test.k8s.elastic.co/v1` | | |
| `kind` _string_ | `Embedded` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `a` _string_ |  |  |  |
| `x` _string_ |  |  |  |
| `value` _[JSON](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#json-v1-apiextensions-k8s-io)_ |  |  |  |


#### <a id="github-com-elastic-crd-ref-docs-api-v1-embedded1">Embedded1</a>







_Appears in:_
- [Embedded](#github-com-elastic-crd-ref-docs-api-v1-embedded)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `x` _string_ |  |  |  |
| `value` _[JSON](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#json-v1-apiextensions-k8s-io)_ |  |  |  |


#### <a id="github-com-elastic-crd-ref-docs-api-v1-embeddedx">EmbeddedX</a>







_Appears in:_
- [Embedded](#github-com-elastic-crd-ref-docs-api-v1-embedded)
- [Embedded1](#github-com-elastic-crd-ref-docs-api-v1-embedded1)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `x` _string_ |  |  |  |
| `value` _[JSON](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#json-v1-apiextensions-k8s-io)_ |  |  |  |


#### <a id="github-com-elastic-crd-ref-docs-api-v1-guestbook">Guestbook</a>



Guestbook is the Schema for the guestbooks API.



_Appears in:_
- [GuestbookList](#github-com-elastic-crd-ref-docs-api-v1-guestbooklist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `webapp.test.k8s.elastic.co/v1` | | |
| `kind` _string_ | `Guestbook` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[GuestbookSpec](#github-com-elastic-crd-ref-docs-api-v1-guestbookspec)_ |  | \{ page:1 \} |  |


#### <a id="github-com-elastic-crd-ref-docs-api-v1-guestbookentry">GuestbookEntry</a>



GuestbookEntry defines an entry in a guest book.



_Appears in:_
- [GuestbookSpec](#github-com-elastic-crd-ref-docs-api-v1-guestbookspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `name` _string_ | Name of the guest (pipe \| should be escaped) |  | MaxLength: 80 <br />Pattern: `0*[a-z0-9]*[a-z]*[0-9]` <br />Required: \{\} <br /> |
| `tags` _string array_ | Tags of the entry. |  | items:Pattern: `[a-z]*` <br /> |
| `time` _[Time](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#time-v1-meta)_ | Time of entry |  |  |
| `comment` _string_ | Comment by guest. This can be a multi-line comment.<br />Like this one.<br />Now let's test a list:<br />* a<br />* b<br />Another isolated comment.<br />Looks good? |  | Pattern: `0*[a-z0-9]*[a-z]*[0-9]*\|\s` <br /> |
| `rating` _[Rating](#github-com-elastic-crd-ref-docs-api-v1-rating)_ | Rating provided by the guest |  | Maximum: 5 <br />Minimum: 1 <br /> |
| `email` _string_ | Email is the email address of the guest (required field using +required marker) |  | Required: \{\} <br /> |
| `location` _string_ | Location is the location of the guest (required field using +k8s:required marker) |  | Required: \{\} <br /> |
| `phone` _string_ | Phone is the phone number of the guest (optional field using +optional marker) |  | Optional: \{\} <br /> |
| `company` _string_ | Company is the company of the guest (optional field using +k8s:optional marker) |  | Optional: \{\} <br /> |


#### <a id="github-com-elastic-crd-ref-docs-api-v1-guestbookheader">GuestbookHeader</a>

_Underlying type:_ _string_

GuestbookHeaders are strings to include at the top of a page.



_Appears in:_
- [GuestbookSpec](#github-com-elastic-crd-ref-docs-api-v1-guestbookspec)



#### <a id="github-com-elastic-crd-ref-docs-api-v1-guestbooklist">GuestbookList</a>



GuestbookList contains a list of Guestbook.





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `webapp.test.k8s.elastic.co/v1` | | |
| `kind` _string_ | `GuestbookList` | | |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Guestbook](#github-com-elastic-crd-ref-docs-api-v1-guestbook) array_ |  |  |  |


#### <a id="github-com-elastic-crd-ref-docs-api-v1-guestbookspec">GuestbookSpec</a>



GuestbookSpec defines the desired state of Guestbook.



_Appears in:_
- [Guestbook](#github-com-elastic-crd-ref-docs-api-v1-guestbook)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `page` _[PositiveInt](#github-com-elastic-crd-ref-docs-api-v1-positiveint)_ | Page indicates the page number | 1 | Minimum: 1 <br /> |
| `entries` _[GuestbookEntry](#github-com-elastic-crd-ref-docs-api-v1-guestbookentry) array_ | Entries contain guest book entries for the page |  |  |
| `selector` _[LabelSelector](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#labelselector-v1-meta)_ | Selector selects something |  |  |
| `headers` _[GuestbookHeader](#github-com-elastic-crd-ref-docs-api-v1-guestbookheader) array_ | Headers contains a list of header items to include in the page |  | MaxItems: 10 <br />UniqueItems: true <br /> |
| `certificateRef` _[SecretObjectReference](https://gateway-api.sigs.k8s.io/references/spec/#gateway.networking.k8s.io/v1beta1.SecretObjectReference)_ | CertificateRef is a reference to a secret containing a certificate |  |  |
| `str` _[CommonString](#github-com-elastic-crd-ref-docs-api-common-commonstring)_ |  |  |  |
| `enum` _[MyEnum](#github-com-elastic-crd-ref-docs-api-v1-myenum)_ | Enumeration is an example of an aliased enumeration type |  | Enum: [MyFirstValue MySecondValue] <br /> |
| `digest` _string_ | Digest is the content-addressable identifier of the guestbook |  | Pattern: `^sha256:[a-fA-F0-9]\{64\}$` <br /> |




#### <a id="github-com-elastic-crd-ref-docs-api-v1-myenum">MyEnum</a>

_Underlying type:_ _string_



_Validation:_
- Enum: [MyFirstValue MySecondValue]

_Appears in:_
- [GuestbookSpec](#github-com-elastic-crd-ref-docs-api-v1-guestbookspec)

| Field | Description |
| `MyFirstValue` | MyFirstValue is an interesting value to use<br /> |
| `MySecondValue` | MySecondValue is what you use when you can't use MyFirstValue<br /> |


#### <a id="github-com-elastic-crd-ref-docs-api-v1-positiveint">PositiveInt</a>

_Underlying type:_ _integer_



_Validation:_
- Minimum: 1

_Appears in:_
- [GuestbookSpec](#github-com-elastic-crd-ref-docs-api-v1-guestbookspec)



#### <a id="github-com-elastic-crd-ref-docs-api-v1-rating">Rating</a>

_Underlying type:_ _integer_

Rating is the rating provided by a guest.

_Validation:_
- Maximum: 5
- Minimum: 1

_Appears in:_
- [GuestbookEntry](#github-com-elastic-crd-ref-docs-api-v1-guestbookentry)





#### <a id="github-com-elastic-crd-ref-docs-api-v1-underlying">Underlying</a>



Underlying tests that Underlying1's underlying type is Underlying2 instead of string.





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `webapp.test.k8s.elastic.co/v1` | | |
| `kind` _string_ | `Underlying` | | |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `a` _[Underlying1](#github-com-elastic-crd-ref-docs-api-v1-underlying1)_ |  | b | MaxLength: 10 <br /> |


#### <a id="github-com-elastic-crd-ref-docs-api-v1-underlying1">Underlying1</a>

_Underlying type:_ _[Underlying2](#github-com-elastic-crd-ref-docs-api-v1-underlying2)_

Underlying1 has an underlying type with an underlying type

_Validation:_
- MaxLength: 10

_Appears in:_
- [Underlying](#github-com-elastic-crd-ref-docs-api-v1-underlying)



#### <a id="github-com-elastic-crd-ref-docs-api-v1-underlying2">Underlying2</a>

_Underlying type:_ _string_

Underlying2 is a string alias

_Validation:_
- MaxLength: 10

_Appears in:_
- [Underlying1](#github-com-elastic-crd-ref-docs-api-v1-underlying1)



