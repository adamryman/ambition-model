
<style type="text/css">

body{
    font-family      : helvetica, arial, freesans, clean, sans-serif;
    color            : #003269;
    background-color : #fff;
    border-color     : #999999;
    border-width     : 2px;
    line-height      : 1.5;
    margin           : 2em 3em;
    text-align       :left;
    font-size        : 16px;
    padding          : 0 100px 0 100px;

    width         : 1024px;
    margin-top    : 0px;
    margin-bottom : 2em;
    margin-left   : auto;
    margin-right  : auto;
}

h1 {
    font-family : 'Gill Sans Bold', 'Optima Bold', Arial, sans-serif;
    color       : #577AD3;
    font-weight : 400;
    font-size   : 48px;
}
h2{
    margin-bottom : 1em;
    padding-top   : 0.5em;
    color         : #003269;
    font-size     : 36px;
}
h3{
    border-bottom : 1px dotted #aaa;
    color         : #4660A4;
    font-size     : 30px;
}
h4 {
    font-size: 24px;
}
h5 {
    font-size: 18px;
}
code {
    font-family      : Consolas, "Inconsolata", Menlo, Monaco, Lucida Console, Liberation Mono, DejaVu Sans Mono, Bitstream Vera Sans Mono, Courier New, monospace, serif; /* Taken from the stackOverflow CSS*/
    background-color : #f5f5f5;
    border           : 1px solid #e1e1e8;
}


pre {
    display          : block;
    background-color : #f5f5f5;
    border           : 1px solid #ccc;
    padding          : 3px 3px 3px 3px;
}
pre code {
    white-space      : pre-wrap;
    padding          : 0;
    border           : 0;
    background-color : code;
}

table {
	border-collapse: collapse; border-spacing: 0;
	width: 100%;
	margin-bottom : 3em;
}
td, th {
	vertical-align: top;
	padding: 4px 10px;
	border: 1px solid #9BC3EB;
}
tr:nth-child(even) td, tr:nth-child(even) th {
	background: #EBF4FE;
}
th:nth-child(4) {
	width: auto;
}

</style>
# ambition

## ambition.proto

### Messages

<a name="ReadActionsRequest"></a>

#### ReadActionsRequest

| Name | Type | Field Number | Description|
| ---- | ---- | ------------ | -----------|
| UserId | TYPE_INT64 | 1 |  |

<a name="ActionResponse"></a>

#### ActionResponse

| Name | Type | Field Number | Description|
| ---- | ---- | ------------ | -----------|
| Action | [Action](#Action) | 1 |  |
| Error | TYPE_STRING | 2 |  |

<a name="ActionsResponse"></a>

#### ActionsResponse

| Name | Type | Field Number | Description|
| ---- | ---- | ------------ | -----------|
| Actions | [Action](#Action) | 1 |  |
| Error | TYPE_STRING | 4 |  |

<a name="Action"></a>

#### Action

| Name | Type | Field Number | Description|
| ---- | ---- | ------------ | -----------|
| ActionId | TYPE_INT64 | 1 |  |
| UserId | TYPE_INT64 | 2 |  |
| ActionName | TYPE_STRING | 3 |  |
| TrelloId | TYPE_STRING | 4 |  |

<a name="CreateOccurrenceRequest"></a>

#### CreateOccurrenceRequest

| Name | Type | Field Number | Description|
| ---- | ---- | ------------ | -----------|
| ActionId | TYPE_INT64 | 1 |  |
| Datetime | TYPE_STRING | 2 |  |
| Data | TYPE_STRING | 3 |  |

<a name="ReadOccurrencesRequest"></a>

#### ReadOccurrencesRequest

| Name | Type | Field Number | Description|
| ---- | ---- | ------------ | -----------|
| UserId | TYPE_INT64 | 1 |  |
| ActionId | TYPE_INT64 | 2 |  |

<a name="ReadOccurrenceRequest"></a>

#### ReadOccurrenceRequest

| Name | Type | Field Number | Description|
| ---- | ---- | ------------ | -----------|
| OccurrenceId | TYPE_INT64 | 1 |  |

<a name="OccurrenceResponse"></a>

#### OccurrenceResponse

| Name | Type | Field Number | Description|
| ---- | ---- | ------------ | -----------|
| Occurrence | [Occurrence](#Occurrence) | 1 |  |
| Error | TYPE_STRING | 2 |  |

<a name="OccurrencesResponse"></a>

#### OccurrencesResponse

| Name | Type | Field Number | Description|
| ---- | ---- | ------------ | -----------|
| Occurrences | [Occurrence](#Occurrence) | 1 |  |
| Error | TYPE_STRING | 2 |  |

<a name="Occurrence"></a>

#### Occurrence

| Name | Type | Field Number | Description|
| ---- | ---- | ------------ | -----------|
| OccurrenceId | TYPE_INT64 | 1 |  |
| ActionId | TYPE_INT64 | 2 |  |
| Datetime | TYPE_STRING | 3 |  |
| Data | TYPE_STRING | 4 |  |

### Services

#### AmbitionService

| Method Name | Request Type | Response Type | Description|
| ---- | ---- | ------------ | -----------|
| ReadActions | ReadActionsRequest | ActionsResponse |  |
| ReadAction | Action | ActionResponse |  |
| CreateAction | Action | ActionResponse |  |
| ReadOccurrences | Occurrence | OccurrenceResponse |  |
| CreateOccurrence | Occurrence | OccurrenceResponse |  |

#### AmbitionService - Http Methods

