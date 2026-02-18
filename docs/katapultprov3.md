[Original Source](https://raw.githubusercontent.com/KatapultDevelopment/katapult-pro-api-documentation/refs/heads/main/v3/README.md)

<!-- NOTE: THIS FILE IS COPIED FROM OUR INTERNAL DESIGN DOCUMENTATION AND SHOULD NOT BE MODIFIED HERE -->
<!--
TODOs:
- Add descriptions to calls
- Add resource definitions (for responses)
- Add examples?
-->

# Requests

## Jobs

Summary of endpoints:

```sh
GET /v3/jobs
GET /v3/jobs/:job_id
POST /v3/jobs
POST /v3/jobs/:job_id
POST /v3/jobs/:job_id/status
```

<!-- TODO (04/28/25): Add these back when the calls are complete
  POST /v3/jobs/:job_id/duplicate
  POST /v3/jobs/:job_id/transfer
-->

### List all jobs

```sh
GET https://katapultpro.com/api/v3/jobs
```

| Query Parameter   | Type     | Description                                                                                                 |
| ----------------- | -------- | ----------------------------------------------------------------------------------------------------------- |
| `includeArchived` | `string` | If "true", archived jobs will be included in the results. By default, only non-archived jobs are included.  |
| `metadataFilter`  | `string` | Metadata filter for the job list (structure as `{attribute}:{value}`, and separate multiples with a comma). |

Gets a list of jobs accessible to the requester. The list entries do not contain
the full job data.

### Get a job (partial data)

```sh
GET https://katapultpro.com/api/v3/jobs/:job_id
```

Route Parameters:

- `job_id`: Id of the job.

| Query Parameter | Type     | Description                                                                                                                                                                                                               |
| --------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `paths`         | `string` | Comma-separated list of data paths to return in the response. The following paths are allowed: `name \| job_creator \| job_owner \| project_folder \| project_id \| status \| done \| map_styles \| metadata \| sharing`. |

Gets partial job data for the specified job.

### Create a job

```sh
POST https://katapultpro.com/api/v3/jobs
```

| Body Field   | Type     | Required | Description                                                                                               |
| ------------ | -------- | :------: | --------------------------------------------------------------------------------------------------------- |
| `name`       | `string` |    ✓     | Name of the job.                                                                                          |
| `model`      | `string` |    ✓     | Model of the job. This value is stored on the `job_creator` property.                                     |
| `map_styles` | `string` |          | Map styles for the job.                                                                                   |
| `metadata`   | `object` |          | Metadata for the job. Must be formatted as a flat map.                                                    |
| `sharing`    | `object` |          | Sharing settings for the job. Must be formatted as a flat map. Owner company will be automatically added. |

Creates a new job using the provided data.

### Update a job

```sh
POST https://katapultpro.com/api/v3/jobs/:job_id
```

Route Parameters:

- `job_id`: Id of the job. The specified job must exist.

| Body Field   | Type     | Required | Description                                                           |
| ------------ | -------- | :------: | --------------------------------------------------------------------- |
| `name`       | `string` |          | Name of the job.                                                      |
| `model`      | `string` |          | Model of the job. This value is stored on the `job_creator` property. |
| `map_styles` | `string` |          | Map styles for the job.                                               |
| `metadata`   | `object` |          | Metadata for the job. Must be formatted as a flat map.                |
| `sharing`    | `object` |          | Sharing settings for the job. Must be formatted as a flat map.        |

Updates the specified job with the provided data.

### Get job status

```sh
GET https://katapultpro.com/api/v3/jobs/:job_id/status
```

Route Parameters:

- `job_id`: Id of the job. The specified job must exist.

Gets the status of the specified job, either `active` or `archived`.

### Update job status

```sh
POST https://katapultpro.com/api/v3/jobs/:job_id/status
```

Route Parameters:

- `job_id`: Id of the job. The specified job must exist.

| Body Field | Type                     | Required | Description         |
| ---------- | ------------------------ | :------: | ------------------- |
| `status`   | `'active' \| 'archived'` |    ✓     | Status for the job. |

Updates the status of the specified job.

<!-- TODO (04/28/25): Add these back when the calls are complete
### Duplicate a job
```sh
POST https://katapultpro.com/api/v3/jobs/:job_id/duplicate
```

Route Parameters:
- `job_id`: Id of the job to duplicate. The specified job must exist.

| Body Field | Type | Required | Description |
| --- | --- | :---: | --- |
| `name` | `string` | | Name of the new job. If not specified, a new name will be generated: `{Source Job Name} (copy)` |
| `transfer_to` | `string` | | Id of a company to transfer ownership of the duplicated job to. If not specified, no transfer will occur. |

### Transfer a job
```sh
POST https://katapultpro.com/api/v3/jobs/:job_id/transfer
```

Route Parameters:
- `job_id`: Id of the job to transfer. The specified job must exist.

| Body Field | Type | Required | Description |
| --- | --- | :---: | --- |
| `transfer_to` | `string` | ✓ | Id of a company to transfer ownership of the job to. |
-->

## Nodes

Summary of endpoints:

```sh
GET /v3/jobs/:job_id/nodes
GET /v3/jobs/:job_id/nodes/:node_id
POST /v3/jobs/:job_id/nodes
POST /v3/jobs/:job_id/nodes/:node_id
POST /v3/jobs/:job_id/nodes/:node_id/photos
DELETE /v3/jobs/:job_id/nodes/:node_id
```

### Get all nodes

```sh
GET https://katapultpro.com/api/v3/jobs/:job_id/nodes
```

Route Parameters:

- `job_id`: Id of the job.

Gets all nodes in the specified job.

### Get a node

```sh
GET https://katapultpro.com/api/v3/jobs/:job_id/nodes/:node_id
```

Route Parameters:

- `job_id`: Id of the job.
- `node_id`: Id of the node.

Gets the specified node.

### Create a node

```sh
POST https://katapultpro.com/api/v3/jobs/{job_id}/nodes
```

Route Parameters:

- `job_id`: Id of the job to create the node in.

| Body Field       | Type     | Required | Description                                                                                                            |
| ---------------- | -------- | :------: | ---------------------------------------------------------------------------------------------------------------------- |
| `latitude`       | `number` |    ✓     | Latitude of the node.                                                                                                  |
| `longitude`      | `number` |    ✓     | Longitude of the node.                                                                                                 |
| `attributes`     | `object` |          | Full attributes object for the node. Must be formatted as an [entity attribute list](#nodes-connections-and-sections). |
| `add_attributes` | `object` |          | Attributes to add to the node (instance ids will be created automatically). Must be formatted as a flat map.           |

Creates a new node using the provided data.

### Update a node

```sh
POST https://katapultpro.com/api/v3/jobs/:job_id/nodes/:node_id
```

Route Parameters:

- `job_id`: Id of the job.
- `node_id`: Id of the node. If the node does not exist, a new node will be
  created with the specified node id (provided it meets the
  [id-requirements](#resource-id-requirements)).

| Query Parameter | Type     | Description                                                                                                 |
| --------------- | -------- | ----------------------------------------------------------------------------------------------------------- |
| `onlyIfExists`  | `string` | If "true", the node will only be updated if it exists, instead of being created with the specified node id. |

| Body Field          | Type     | Required | Description                                                                                                            |
| ------------------- | -------- | :------: | ---------------------------------------------------------------------------------------------------------------------- |
| `latitude`          | `number` |          | Latitude for the node.                                                                                                 |
| `longitude`         | `number` |          | Longitude for the node.                                                                                                |
| `remove_attributes` | `array`  |          | Attributes to remove (all instances) from the node. Must be formatted as a list.                                       |
| `attributes`        | `object` |          | Full attributes object for the node. Must be formatted as an [entity attribute list](#nodes-connections-and-sections). |
| `add_attributes`    | `object` |          | Attributes to add to the node (instance ids will be created automatically). Must be formatted as a flat map.           |

Updates the specified node with the provided data.

### Upload a photo to a node

```sh
POST https://katapultpro.com/api/v3/jobs/:job_id/nodes/:node_id/photos
```

Route Parameters:

- `job_id`: Id of the job to upload the photo to.
- `node_id`: Id of the node to upload the photo to.

| Query Parameter     | Type     | Description                                                                                                              |
| ------------------- | -------- | ------------------------------------------------------------------------------------------------------------------------ |
| `association_value` | `string` | Determines how the photo will be associated to the node. Valid values are "main" and "true". The default value is "true" |

Uploads a photo (must be an `image/jpeg` file) and associates it to the
specified node.

### Delete a node

```sh
DELETE https://katapultpro.com/api/v3/jobs/:job_id/nodes/:node_id
```

Route Parameters:

- `job_id`: Id of the job.
- `node_id`: Id of the node.

Deletes the specified node.

## Connections

Summary of endpoints:

```sh
GET /v3/jobs/:job_id/connections
GET /v3/jobs/:job_id/connections/:connection_id
POST /v3/jobs/:job_id/connections
POST /v3/jobs/:job_id/connections/:connection_id
DELETE /v3/jobs/:job_id/connections/:connection_id
```

### Get all connections

```sh
GET https://katapultpro.com/api/v3/jobs/:job_id/connections
```

Route Parameters:

- `job_id`: Id of the job.

Gets all connections (and their sections) in the specified job.

### Get a connection

```sh
GET https://katapultpro.com/api/v3/jobs/:job_id/connections/:connection_id
```

Route Parameters:

- `job_id`: Id of the job.
- `connection_id`: Id of the connection.

Gets the specified connection (and its sections).

### Create a connection

```sh
POST https://katapultpro.com/api/v3/jobs/:job_id/connections
```

Route Parameters:

- `job_id`: Id of the job to create the connection in.

| Body Field       | Type     | Required | Description                                                                                                                  |
| ---------------- | -------- | :------: | ---------------------------------------------------------------------------------------------------------------------------- |
| `node_id_1`      | `string` |    ✓     | Id of the first node.                                                                                                        |
| `node_id_2`      | `string` |    ✓     | Id of the second node.                                                                                                       |
| `attributes`     | `object` |          | Full attributes object for the connection. Must be formatted as an [entity attribute list](#nodes-connections-and-sections). |
| `add_attributes` | `object` |          | Attributes to add to the connection (instance ids will be created automatically). Must be formatted as a flat map.           |

Creates a new connection using the provided data.

### Update a connection

```sh
POST https://katapultpro.com/api/v3/jobs/:job_id/connections/:connection_id
```

Route Parameters:

- `job_id`: Id of the job.
- `connection_id`: Id of the connection. If the connection does not exist, a new
  connection will be created with the specified connection id (provided it meets
  the [id-requirements](#resource-id-requirements)).

| Query Parameter | Type     | Description                                                                                                             |
| --------------- | -------- | ----------------------------------------------------------------------------------------------------------------------- |
| `onlyIfExists`  | `string` | If "true", the connection will only be updated if it exists, instead of being created with the specified connection id. |

| Body Field          | Type     | Required | Description                                                                                                                  |
| ------------------- | -------- | :------: | ---------------------------------------------------------------------------------------------------------------------------- |
| `node_id_1`         | `string` |          | Id of the first node.                                                                                                        |
| `node_id_2`         | `string` |          | Id of the second node.                                                                                                       |
| `remove_attributes` | `array`  |          | Attributes to remove (all instances) from the connection. Must be formatted as a list.                                       |
| `attributes`        | `object` |          | Full attributes object for the connection. Must be formatted as an [entity attribute list](#nodes-connections-and-sections). |
| `add_attributes`    | `object` |          | Attributes to add to the connection (instance ids will be created automatically). Must be formatted as a flat map.           |

Updates the specified connection with the provided data.

### Delete a connection

```sh
DELETE https://katapultpro.com/api/v3/jobs/:job_id/connections/:connection_id
```

Route Parameters:

- `job_id`: Id of the job.
- `connection_id`: Id of the connection.

Deletes the specified connection (and all of its sections).

## Sections

Summary of endpoints:

```sh
GET /v3/jobs/:job_id/connections/:connection_id/sections
GET /v3/jobs/:job_id/connections/:connection_id/sections/:section_id
POST /v3/jobs/:job_id/connections/:connection_id/sections
POST /v3/jobs/:job_id/connections/:connection_id/sections/:section_id
POST /v3/jobs/:job_id/connections/:connection_id/sections/:section_id/photos
DELETE /v3/jobs/:job_id/connections/:connection_id/sections/:section_id
```

### Get all sections on a connection

```sh
GET https://katapultpro.com/api/v3/jobs/:job_id/connections/:connection_id/sections
```

Route Parameters:

- `job_id`: Id of the job.
- `connection_id`: Id of the connection.

Gets all sections on the specified connection.

### Get a section

```sh
GET https://katapultpro.com/api/v3/jobs/:job_id/connections/:connection_id/sections/:section_key
```

Route Parameters:

- `job_id`: Id of the job.
- `connection_id`: Id of the connection.
- `section_key`: Key of the section.

Gets the specified section.

### Create a section

```sh
POST https://katapultpro.com/api/v3/jobs/:job_id/connections/:connection_id/sections
```

Route Parameters:

- `job_id`: Id of the job.
- `connection_id`: Id of the connection to create the section in.

| Body Field       | Type      | Required | Description                                                                                                               |
| ---------------- | --------- | :------: | ------------------------------------------------------------------------------------------------------------------------- |
| `make_midpoint`  | `boolean` |          | If true, the section will be created as a midpoint section.                                                               |
| `latitude`       | `number`  |          | Latitude of the section.                                                                                                  |
| `longitude`      | `number`  |          | Longitude of the section.                                                                                                 |
| `attributes`     | `object`  |          | Full attributes object for the section. Must be formatted as an [entity attribute list](#nodes-connections-and-sections). |
| `add_attributes` | `object`  |          | Attributes to add to the section (instance ids will be created automatically). Must be formatted as a flat map.           |

Creates a new section using the provided data.

### Update a section

```sh
POST https://katapultpro.com/api/v3/jobs/:job_id/connections/:connection_id/sections/:section_key
```

Route Parameters:

- `job_id`: Id of the job.
- `connection_id`: Id of the connection.
- `section_key`: Key of the section. If the section does not exist, a new
  section will be created with the specified section id (provided it meets the
  [id-requirements](#resource-id-requirements)).

| Query Parameter | Type     | Description                                                                                                       |
| --------------- | -------- | ----------------------------------------------------------------------------------------------------------------- |
| `onlyIfExists`  | `string` | If "true", the section will only be updated if it exists, instead of being created with the specified section id. |

| Body Field          | Type     | Required | Description                                                                                                               |
| ------------------- | -------- | :------: | ------------------------------------------------------------------------------------------------------------------------- |
| `latitude`          | `number` |          | Latitude for the section.                                                                                                 |
| `longitude`         | `number` |          | Longitude for the section.                                                                                                |
| `remove_attributes` | `array`  |          | Attributes to remove (all instances) from the section. Must be formatted as a list.                                       |
| `attributes`        | `object` |          | Full attributes object for the section. Must be formatted as an [entity attribute list](#nodes-connections-and-sections). |
| `add_attributes`    | `object` |          | Attributes to add to the section (instance ids will be created automatically). Must be formatted as a flat map.           |

Updates the specified section with the provided data.

### Upload a photo to a section

```sh
POST https://katapultpro.com/api/v3/jobs/:job_id/connections/:connection_id/sections/:section_id/photos
```

Route Parameters:

- `job_id`: Id of the job to upload the photo to.
- `connection_id`: Id of the connection that the section is on.
- `section_id`: Id of the section to upload the photo to.

| Query Parameter     | Type     | Description                                                                                                                 |
| ------------------- | -------- | --------------------------------------------------------------------------------------------------------------------------- |
| `association_value` | `string` | Determines how the photo will be associated to the section. Valid values are "main" and "true". The default value is "true" |

Uploads a photo (must be an `image/jpeg` file) and associates it to the
specified section.

### Delete a section

```sh
DELETE https://katapultpro.com/api/v3/jobs/:job_id/connections/:connection_id/sections/:section_key
```

Route Parameters:

- `job_id`: Id of the job.
- `connection_id`: Id of the connection.
- `section_key`: Key of the section.

Deletes the specified section.

## Photos

Summary of endpoints:

```sh
GET /v3/jobs/:job_id/photos
GET /v3/jobs/:job_id/photos/:photo_id
POST /v3/jobs/:job_id/photos
POST /v3/jobs/:job_id/photos/:photo_id/associate
```

### Get all photos

```sh
GET https://katapultpro.com/api/v3/jobs/:job_id/photos
```

Route Parameters:

- `job_id`: Id of the job.

Gets all photo records in the specified job.

### Get a photo

```sh
GET https://katapultpro.com/api/v3/jobs/:job_id/photos/:photo_id
```

Route Parameters:

- `job_id`: Id of the job.
- `photo_id`: Id of the photo.

Gets the specified photo record.

### Upload a photo

```sh
POST https://katapultpro.com/api/v3/jobs/:job_id/photos
```

Route Parameters:

- `job_id`: Id of the job to upload the photo to.

Uploads a photo (must be an `image/jpeg` file) to the specified job.

### Associate a photo to an item

```sh
POST https://katapultpro.com/api/v3/jobs/:job_id/photos/:photo_id/associate
```

Route Parameters:

- `job_id`: Id of the job to associate a photo for.
- `photo_id`: Id of the photo to associate

| Body Field          | Type                     | Required | Description                                                                                                                                              |
| ------------------- | ------------------------ | :------: | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `node_id`           | `string`                 |    ✓†    | Id of the node to associate the photo to.                                                                                                                |
| `section_id`        | `string`                 |    ✓‡    | Id of the section to associate the photo to.                                                                                                             |
| `connection_id`     | `string`                 |    ✓‡    | Id of the connection that the section is on.                                                                                                             |
| `association_value` | `'main' \| true \| null` |    ✓     | Determines how the photo will be associated to the item. Valid values are "main", true, and null. If null, the photo will be unassociated from the item. |

† To associate to a node, provide `node_id` and omit `section_id` and
`connection_id`.

‡ To associate to a section, provide `section_id` and `connection_id` and omit
`node_id`.

Associates the specified photo to (or unassociates it from) a node or section.

## Photo Elements

Note: photo element updates via the API currently do _not_ update effective
moves in midspans.

Summary of endpoints:

```sh
GET /v3/jobs/:job_id/photos/:photo_id/photo_elements
GET /v3/jobs/:job_id/photos/:photo_id/photo_elements/:element_id
POST /v3/jobs/:job_id/photos/:photo_id/photo_elements
POST /v3/jobs/:job_id/photos/:photo_id/photo_elements/:element_id
DELETE /v3/jobs/:job_id/photos/:photo_id/photo_elements/:element_id
```

### Get all elements on a photo

```sh
GET https://katapultpro.com/api/v3/jobs/:job_id/photos/:photo_id/photo_elements
```

Route Parameters:

- `job_id`: Id of the job.
- `photo_id`: Id of the photo.

Gets all elements on the specified photo.

### Get a photo element

```sh
GET https://katapultpro.com/api/v3/jobs/:job_id/photos/:photo_id/photo_elements/:element_id
```

Route Parameters:

- `job_id`: Id of the job.
- `photo_id`: Id of the photo.
- `element_id`: Id of the element.

Gets the specified photo element.

### Create a photo element

```sh
POST https://katapultpro.com/api/v3/jobs/:job_id/photos/:photo_id/photo_elements
```

Route Parameters:

- `job_id`: Id of the job.
- `photo_id`: Id of the photo to create the element in.

| Body Field        | Type                                     | Required | Description                                                                                     |
| ----------------- | ---------------------------------------- | :------: | ----------------------------------------------------------------------------------------------- |
| `element_type`    | `string`                                 |    ✓     | Type of the element.                                                                            |
| `pixel_selection` | `{ percentX: number, percentY: number }` |          | Pixel selection for the element.                                                                |
| `manual_height`   | `string`                                 |          | Manual height for the element, in feet-inches notation (i.e. `<feet>-<inches>`).                |
| `attributes`      | `object`                                 |          | Attributes for the element. Must be formatted as a flat map.                                    |
| `parent_id`       | `string`                                 |          | Id of the element to make this element a child of. If omitted, this element will not be nested. |
| `trace_id`        | `string`                                 |          | Id of the trace to add this element to.                                                         |

Creates a new photo element using the provided data.

### Update a photo element

```sh
POST https://katapultpro.com/api/v3/jobs/:job_id/photos/:photo_id/photo_elements/:element_id
```

Route Parameters:

- `job_id`: Id of the job.
- `photo_id`: Id of the photo.
- `element_id`: Id of the element to update. If an element is not found with the
  specified ID, a new element with a random ID will be created.

| Query Parameter | Type     | Description                                                                                                       |
| --------------- | -------- | ----------------------------------------------------------------------------------------------------------------- |
| `onlyIfExists`  | `string` | If "true", the element will only be updated if it exists, instead of being created with the specified element id. |

| Body Field        | Type                                     | Required | Description                                                                            |
| ----------------- | ---------------------------------------- | :------: | -------------------------------------------------------------------------------------- |
| `element_type`    | `string`                                 |          | Type of the element. Can only be set if the element does not already exist.            |
| `pixel_selection` | `{ percentX: number, percentY: number }` |          | Pixel selection for the element.                                                       |
| `manual_height`   | `string`                                 |          | Manual height for the element, in feet-inches notation (i.e. `<feet>-<inches>`).       |
| `attributes`      | `object`                                 |          | Attributes for the element. Must be formatted as a flat map.                           |
| `parent_id`       | `string \| null`                         |          | Id of the element to make this element a child of. Set to null to de-nest the element. |
| `trace_id`        | `string`                                 |          | Id of the trace to add this element to.                                                |

Updates the specified photo element.

### Delete a photo element

```sh
DELETE https://katapultpro.com/api/v3/jobs/:job_id/photos/:photo_id/photo_elements/:element_id
```

Route Parameters:

- `job_id`: Id of the job.
- `photo_id`: Id of the photo.
- `element_id`: Id of the element.

Deletes the specified photo element.

## Photo Calibration Anchors

Note: adding or modifying anchor calibration points via the API will always
un-calibrate the photo (i.e. remove the `stick_align`). Viewing the photo in
Katapult Pro will recalibrate the photo based on the updated anchor calibration
points.

Summary of endpoints:

```sh
GET /v3/jobs/:job_id/photos/:photo_id/calibration_anchors
GET /v3/jobs/:job_id/photos/:photo_id/calibration_anchors/:anchor_id
POST /v3/jobs/:job_id/photos/:photo_id/calibration_anchors
POST /v3/jobs/:job_id/photos/:photo_id/calibration_anchors/:anchor_id
DELETE /v3/jobs/:job_id/photos/:photo_id/calibration_anchors/:anchor_id
```

### Get all calibration anchors on a photo

```sh
GET https://katapultpro.com/api/v3/jobs/:job_id/photos/:photo_id/calibration_anchors
```

Route Parameters:

- `job_id`: Id of the job.
- `photo_id`: Id of the photo.

Gets all calibration anchors on the specified photo.

### Get a photo calibration anchor

```sh
GET https://katapultpro.com/api/v3/jobs/:job_id/photos/:photo_id/calibration_anchors/:anchor_id
```

Route Parameters:

- `job_id`: Id of the job.
- `photo_id`: Id of the photo.
- `anchor_id`: Id of the calibration anchor.

Gets the specified photo calibration anchor.

### Create a photo calibration anchor

```sh
POST https://katapultpro.com/api/v3/jobs/:job_id/photos/:photo_id/calibration_anchors
```

Route Parameters:

- `job_id`: Id of the job.
- `photo_id`: Id of the photo to create the calibration anchor in.

| Body Field        | Type                                     | Required | Description                                        |
| ----------------- | ---------------------------------------- | :------: | -------------------------------------------------- |
| `pixel_selection` | `{ percentX: number, percentY: number }` |    ✓     | Pixel selection for the calibration anchor.        |
| `height`          | `number`                                 |    ✓     | Height of the calibration anchor, in decimal feet. |

Creates a new photo calibration anchor using the provided data.

### Update a photo calibration anchor

```sh
POST https://katapultpro.com/api/v3/jobs/:job_id/photos/:photo_id/calibration_anchors/:anchor_id
```

Route Parameters:

- `job_id`: Id of the job.
- `photo_id`: Id of the photo.
- `anchor_id`: Id of the calibration anchor. If the calibration anchor does not
  exist, a new calibration anchor will be created with the specified anchor id
  (provided it meets the [id-requirements](#resource-id-requirements)).

| Query Parameter | Type     | Description                                                                                                                 |
| --------------- | -------- | --------------------------------------------------------------------------------------------------------------------------- |
| `onlyIfExists`  | `string` | If "true", the calibration anchor will only be updated if it exists, instead of being created with the specified anchor id. |

| Body Field        | Type                                     | Required | Description                                        |
| ----------------- | ---------------------------------------- | :------: | -------------------------------------------------- |
| `pixel_selection` | `{ percentX: number, percentY: number }` |          | Pixel selection for the calibration anchor.        |
| `height`          | `number`                                 |          | Height of the calibration anchor, in decimal feet. |

Updates the specified photo calibration anchor.

### Delete a photo calibration anchor

```sh
DELETE https://katapultpro.com/api/v3/jobs/:job_id/photos/:photo_id/calibration_anchors/:anchor_id
```

Route Parameters:

- `job_id`: Id of the job.
- `photo_id`: Id of the photo.
- `anchor_id`: Id of the calibration anchor.

Deletes the specified photo calibration anchor.

## Traces

Summary of endpoints:

```sh
GET /v3/jobs/:job_id/traces
GET /v3/jobs/:job_id/traces/:trace_id
POST /v3/jobs/:job_id/traces
POST /v3/jobs/:job_id/traces/:trace_id
DELETE /v3/jobs/:job_id/traces/:trace_id
```

### Get all traces

```sh
GET https://katapultpro.com/api/v3/jobs/:job_id/traces
```

Route Parameters:

- `job_id`: Id of the job.

Gets all traces in the specified job.

### Get a trace

```sh
GET https://katapultpro.com/api/v3/jobs/:job_id/traces/:trace_id
```

Route Parameters:

- `job_id`: Id of the job.
- `trace_id`: Id of the trace.

Gets the specified trace.

### Create a trace

```sh
POST https://katapultpro.com/api/v3/jobs/:job_id/traces
```

Route Parameters:

- `job_id`: Id of the job to create the trace in.

| Body Field   | Type     | Required | Description                                                |
| ------------ | -------- | :------: | ---------------------------------------------------------- |
| `trace_type` | `string` |    ✓     | Type of the trace.                                         |
| `attributes` | `object` |          | Attributes for the trace. Must be formatted as a flat map. |

<!-- TODO (04/28/25): Add this back when the call is complete
 | `markers` | `array` | | Array of marker ids to add to the trace. |
 -->

Creates a new trace using the provided data.

### Update a trace

```sh
POST https://katapultpro.com/api/v3/jobs/:job_id/traces/:trace_id
```

Route Parameters:

- `job_id`: Id of the job.
- `trace_id`: Id of the trace. If the trace does not exist, a new trace will be
  created with the specified trace id (provided it meets the
  [id-requirements](#resource-id-requirements)).

| Query Parameter | Type     | Description                                                                                                   |
| --------------- | -------- | ------------------------------------------------------------------------------------------------------------- |
| `onlyIfExists`  | `string` | If "true", the trace will only be updated if it exists, instead of being created with the specified trace id. |

| Body Field   | Type     | Required | Description                                                |
| ------------ | -------- | :------: | ---------------------------------------------------------- |
| `trace_type` | `string` |          | Type of the trace.                                         |
| `attributes` | `object` |          | Attributes for the trace. Must be formatted as a flat map. |

<!-- TODO (04/28/25): Add this back when the call is complete
 | `markers` | `array` | | Array of marker ids to add to the trace. |
 -->

Updates the specified trace with the provided data.

### Delete a trace

```sh
DELETE https://katapultpro.com/api/v3/jobs/:job_id/traces/:trace_id
```

Route Parameters:

- `job_id`: Id of the job.
- `trace_id`: Id of the trace.

Deletes the specified trace.

<!-- TODO (04/28/25): Add these back when the calls are complete
## Users

Summary of endpoints:
```sh
GET /v3/users
GET /v3/users/:user_id
```

### List all users
```sh
GET https://katapultpro.com/api/v3/users
```

| Query Parameter | Type | Description |
| --- | --- | --- |
| `companyId` | `string` | Id of the company to filter by (by default, the requester's root company is used). |

### Get a user
```sh
GET https://katapultpro.com/api/v3/users/:user_id
```

Route Parameters:
- `user_id`: Id of the user.

## Attribute History

Summary of endpoints:
```sh
GET /v3/attribute_history
GET /v3/attribute_history/:record_id
```

### Query attribute history records
```sh
GET https://katapultpro.com/api/v3/attribute_history
```

| Query Parameter | Type | Description |
| --- | --- | --- |
| `jobId` | `string` | Id of the job. |
| `nodeId` | `string` | Id of the node. |
| `connectionId` | `string` | Id of the connection. |
| `sectionKey` | `string` | Key of the section. |
| `attribute` | `string` | Attribute to filter by. |
| `userId` | `string` | Id of the user. |
| `latestOnly` | `boolean` | If true, only the latest record for each attribute will be returned. |
| `fromAllCompanies` | `boolean` | If true, records from all companies will be returned. Can only be used on owned jobs. |
| `fromDate` | `string` | Starting date for time filter (as `MM/DD/YYYY`). |
| `toDate` | `string` | Ending date for time filter (as `MM/DD/YYYY`). |
| `limit` | `number` | Limit the number of records returned. (Default is 100, max is 1000.) |
| `startAfter` | `string` | Id of the record to start after (use with `limit` for pagination). |

### Get an attribute history record
```sh
GET https://katapultpro.com/api/v3/attribute_history/:record_id
```

Route Parameters:
- `record_id`: Id of the record.

## Tracked Actions

Summary of endpoints:
```sh
GET /v3/tracked_actions
GET /v3/tracked_actions/:record_id
```

### Query tracked actions
```sh
GET https://katapultpro.com/api/v3/tracked_actions
```

| Query Parameter | Type | Description |
| --- | --- | --- |
| `jobId` | `string` | Id of the job. |
| `nodeId` | `string` | Id of the node. |
| `actionId` | `string` | Id of the action. |
| `userId` | `string` | Id of the user. |
| `fromDate` | `string` | Starting date for time filter (as `MM/DD/YYYY`). |
| `toDate` | `string` | Ending date for time filter (as `MM/DD/YYYY`). |
| `limit` | `number` | Limit the number of records returned. (Default is 100, max is 1000.) |
| `startAfter` | `string` | Id of the record to start after (use with `limit` for pagination). |

### Get a tracked action
```sh
GET https://katapultpro.com/api/v3/tracked_actions/:record_id
```

Route Parameters:
- `record_id`: Id of the record.

-->

# Responses

## Success

For individual resources, the data field contains the resource:

```js
type EntityResponse = {
  status: "success",
  data: {
    id: string, // Id of the resource
    // Other resource fields
  },
  meta: {
    token_count: number, // Tokens remaining after the request
    last_refill_time: number, // Epoch time of the last token refill
  }
}
```

For lists of resources, the data field contains an array of resources:

```ts
type ListResponse = {
    status: "success";
    data: [
        {
            id: string; // Id of the list item
            // Other resource fields
        },
        // More resources
    ];
    meta: {
        token_count: number; // Tokens remaining after the request
        last_refill_time: number; // Epoch time of the last token refill
    };
};
```

## Error

```ts
type ErrorResponse = {
    status: "error";
    message: string; // Error message
    type: string; // Type of error in human-readable format (e.g. "not_found")
    meta: {
        token_count: number; // Tokens remaining after the request
        last_refill_time: number; // Epoch time of the last token refill
    };
};
```

# Other API Concepts

## Working with Attributes

### Nodes, Connections, and Sections

Nodes, connections, and sections all have attributes, and their attributes are
stored as an "entity attribute list": a JSON object mapping attribute names to
values on an entity. The structure of the object is as follows:

```ts
type EntityAttributeList = {
    [attribute_name: string]: {
        [instance_id: string]: any; // Value of this attribute instance
    };
};
```

That is, for each attribute in an attribute list, there may be multiple "values"
(stored under an `instance_id`). The following is an example of a node attribute
list:

```json
{
    "node_type": {
        "-OPWRGCw7wC7DA3Lt6SX": "pole"
    },
    "scid": {
        "-OPWRQOwUdDU4X4f_-8u": "001"
    },
    "note": {
        "-OPWRVKh0JHKTnZrY8O5": "This is a note",
        "-OPWRVYCKdxh34WoQXAv": "This is another note"
    }
}
```

The one-to-many relationship between attribute names and values makes working
with entity attribute lists a bit more complex. As a result, Katapult recommends
only setting one value per attribute when possible. The API also provides some
convenience methods for working with attributes.

To **add** attributes to an entity list without specifying instance ids, you can
use the `add_attributes` field on POST requests. This field is a flat map of
attribute names to values, and the API will automatically create instance ids
for each value.

For example, the following request:

```json
{
    "add_attributes": {
        "node_type": "pole",
        "scid": "001"
    }
}
```

will create new attribute instances for `node_type` and `scid` with the values
"pole" and "001", respectively. The API will automatically generate instance ids
for these attributes.

To **remove** attributes from an entity list without specifying instance ids,
you can use the `remove_attributes` field on POST requests. This field is a list
of attribute names to remove, and the API will remove all instances of those
attributes from the entity. For example, the following request:

```json
{
    "remove_attributes": ["node_type", "scid"]
}
```

will remove _all_ instances of the `node_type` and `scid` attributes from the
entity.

If you need to **update** an attribute value, you must use the `attributes`
field on POST requests. This field is a partial entity attribute list with a
number of behaviors:

- Any attributes not specified in the request will not be modified.
- If an attribute instance id is specified, the value will be updated for that
  instance id.
- Setting either the attribute name or instance id to `null` will remove that
  attribute or instance id from the entity, respectively.

For example, this request:

```json
{
    "attributes": {
        "node_type": {
            "-OPWRGCw7wC7DA3Lt6SX": "reference"
        },
        "scid": null,
        "note": {
            "-OPWRVYCKdxh34WoQXAv": null
        }
    }
}
```

would do the following:

- Update the `node_type` attribute instance (id = `-OPWRGCw7wC7DA3Lt6SX`) to
  "reference".
- Remove the `scid` attribute from the entity (if there were more than one
  instance, all instances would be removed).
- Remove the `note` attribute instance (id = `-OPWRVYCKdxh34WoQXAv`) from the
  entity.

Note that to update specific attribute instances, you must have the instance id.

The `remove_attributes`, `attributes`, and `add_attributes` fields can be used
together in a single request, and they are applied in that order.

### Jobs

Jobs have attributes (under the `metadata` field) that are stored as a flat map;
job attributes may not have more than one value. The structure of the object is
as follows:

```ts
type JobMetadata = {
    [attribute_name: string]: any; // Value of this attribute
};
```

To add, update, or remove attributes from a job, you can use the `metadata`
field on POST requests. This request:

```json
{
    "metadata": {
        "description": "A description of the job",
        "contact_number": null
    }
}
```

would do the following:

- Add (or update) the `description` attribute to "A description of the job".
- Remove the `contact_number` attribute from the job.

### Photo Elements and Traces

Photo elements and traces have their attributes stored as a flat map too and can
be updated in the same way as job attributes. However, photo element and trace
attributes are not stored under a `metadata` or `attributes` field; instead,
they are stored directly on the element or trace object. To access an attribute
you've set on either a photo element or a trace, read the attribute name
property directly from the root of the object.

## Resource ID Requirements

It is possible to create resources with user-defined IDs through the POST update
endpoints. However, there are some requirements for ids that you should be aware
of:

- All ids must be alphanumeric (a-z, A-Z, 0-9) and may contain dashes (-) and
  underscores (_).
- All ids must be between 20 and 256 characters long.
- All ids must be unique within the job.

Attribute instance ids are not subject to the length or uniqueness requirements,
but they must still be alphanumeric and may contain dashes and underscores.

# Rate Limits

All REST APIs employ rate limits to ensure fair use and prevent abuse. Version 3
of Katapult Pro's API uses two different rate limiting strategies to help manage
the volume of requests being processed. If you send too many requests at one
time, you may receive an error response with a status code of `429`.

## General Rate Limit

Each API key is subject to a general rate limit of **1 call per 50
milliseconds**. This limit prevents sudden bursts of calls from overloading
connected APIs.

## Token System

The token system works by assigning a token cost to each call. Generally, calls
with more database transactions will have a higher token cost.

Each individual API key has its own "bucket" of tokens. When a call is made, the
cost of the call is deducted from the caller's token bucket count. Once the
token bucket is depleted, further calls will be blocked until the bucket is
refilled.

The token bucket systems operates under the following set of parameters:

- The default token bucket allotment is **10000 tokens**.
  - Default token allotment is subject to change.
- All token buckets are refilled **every minute**.
- Successful calls deduct the cost of the call from the token bucket.
- Blocked calls will error and return a status code of `429`.
- All API responses (success and error) include a `meta` object that contains
  information about the API key's token bucket state, including `token_count`
  (the number of tokens remaining) and `last_refill_time` (the last time that
  the bucket was refilled)
  - The bucket will be refilled when 60 seconds have elapsed since the
    `last_refill_time`.

Listed below are the current costs of our API calls:

- **POST** and **DELETE** calls: 10 tokens each
- **GET** calls: 1 token each
