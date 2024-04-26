# Table

The tabular result format can be enabled via APLQueryParams.ResultFormat or QueryParams.ResultFormat.


## Fields

| Field                                                                                                                                          | Type                                                                                                                                           | Required                                                                                                                                       | Description                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| `Buckets`                                                                                                                                      | [*components.BucketInfo](../../models/components/bucketinfo.md)                                                                                | :heavy_minus_sign:                                                                                                                             | The standard mode of operation for Kirby is to create buckets on the _time column,                                                             |
| `Columns`                                                                                                                                      | [][][components.Columns](../../models/components/columns.md)                                                                                   | :heavy_minus_sign:                                                                                                                             | Columns contain a series of arrays with the raw result data.<br/>The columns here line up with the fields in the Fields array.                 |
| `Fields`                                                                                                                                       | [][components.FieldInfo](../../models/components/fieldinfo.md)                                                                                 | :heavy_minus_sign:                                                                                                                             | Fields contain information about the fields included in these results.<br/>The order of the fields match up with the order of the data in Columns. |
| `Groups`                                                                                                                                       | [][components.GroupInfo](../../models/components/groupinfo.md)                                                                                 | :heavy_minus_sign:                                                                                                                             | Groups specifies which grouping operations has been performed on the results.                                                                  |
| `Name`                                                                                                                                         | **string*                                                                                                                                      | :heavy_minus_sign:                                                                                                                             | Name is the name assigned to this table. Defaults to "0". The name "_totals" is reserved for system use.                                       |
| `Order`                                                                                                                                        | [][components.Order](../../models/components/order.md)                                                                                         | :heavy_minus_sign:                                                                                                                             | Order echoes the ordering clauses that was used to sort the results.                                                                           |
| `Range`                                                                                                                                        | [*components.RangeInfo](../../models/components/rangeinfo.md)                                                                                  | :heavy_minus_sign:                                                                                                                             | N/A                                                                                                                                            |
| `Sources`                                                                                                                                      | [][components.SourceInfo](../../models/components/sourceinfo.md)                                                                               | :heavy_minus_sign:                                                                                                                             | Sources contain the names of the datasets that contributed data to these results.                                                              |