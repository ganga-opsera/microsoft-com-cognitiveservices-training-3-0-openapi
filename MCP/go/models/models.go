package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// BoundingBox represents the BoundingBox schema from the OpenAPI specification
type BoundingBox struct {
	Height float32 `json:"height"` // Height.
	Left float32 `json:"left"` // Coordinate of the left boundary.
	Top float32 `json:"top"` // Coordinate of the top boundary.
	Width float32 `json:"width"` // Width.
}

// ImageRegionCreateResult represents the ImageRegionCreateResult schema from the OpenAPI specification
type ImageRegionCreateResult struct {
	Regionid string `json:"regionId,omitempty"`
	Top float32 `json:"top"` // Coordinate of the top boundary.
	Width float32 `json:"width"` // Width.
	Imageid string `json:"imageId,omitempty"`
	Tagname string `json:"tagName,omitempty"`
	Created string `json:"created,omitempty"`
	Height float32 `json:"height"` // Height.
	Left float32 `json:"left"` // Coordinate of the left boundary.
	Tagid string `json:"tagId"` // Id of the tag associated with this region.
}

// ImageFileCreateBatch represents the ImageFileCreateBatch schema from the OpenAPI specification
type ImageFileCreateBatch struct {
	Images []ImageFileCreateEntry `json:"images,omitempty"`
	Tagids []string `json:"tagIds,omitempty"`
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Description string `json:"description"` // Gets or sets the description of the tag.
	Id string `json:"id,omitempty"` // Gets the Tag ID.
	Imagecount int `json:"imageCount,omitempty"` // Gets the number of images with this tag.
	Name string `json:"name"` // Gets or sets the name of the tag.
	TypeField string `json:"type"` // Gets or sets the type of the tag.
}

// PredictionQueryTag represents the PredictionQueryTag schema from the OpenAPI specification
type PredictionQueryTag struct {
	Id string `json:"id,omitempty"`
	Maxthreshold float32 `json:"maxThreshold,omitempty"`
	Minthreshold float32 `json:"minThreshold,omitempty"`
}

// ImageUrl represents the ImageUrl schema from the OpenAPI specification
type ImageUrl struct {
	Url string `json:"url"` // Url of the image.
}

// PredictionQueryToken represents the PredictionQueryToken schema from the OpenAPI specification
type PredictionQueryToken struct {
	Orderby string `json:"orderBy,omitempty"`
	Application string `json:"application,omitempty"`
	Continuation string `json:"continuation,omitempty"`
	Starttime string `json:"startTime,omitempty"`
	Iterationid string `json:"iterationId,omitempty"`
	Session string `json:"session,omitempty"`
	Tags []PredictionQueryTag `json:"tags,omitempty"`
	Endtime string `json:"endTime,omitempty"`
	Maxcount int `json:"maxCount,omitempty"`
}

// Project represents the Project schema from the OpenAPI specification
type Project struct {
	Lastmodified string `json:"lastModified,omitempty"` // Gets the date this project was last modified.
	Name string `json:"name"` // Gets or sets the name of the project.
	Settings ProjectSettings `json:"settings"` // Represents settings associated with a project.
	Thumbnailuri string `json:"thumbnailUri,omitempty"` // Gets the thumbnail url representing the image.
	Created string `json:"created,omitempty"` // Gets the date this project was created.
	Description string `json:"description"` // Gets or sets the description of the project.
	Drmodeenabled bool `json:"drModeEnabled,omitempty"` // Gets if the DR mode is on.
	Id string `json:"id,omitempty"` // Gets the project id.
}

// ImageRegionCreateEntry represents the ImageRegionCreateEntry schema from the OpenAPI specification
type ImageRegionCreateEntry struct {
	Top float32 `json:"top"` // Coordinate of the top boundary.
	Width float32 `json:"width"` // Width.
	Height float32 `json:"height"` // Height.
	Imageid string `json:"imageId"` // Id of the image.
	Left float32 `json:"left"` // Coordinate of the left boundary.
	Tagid string `json:"tagId"` // Id of the tag associated with this region.
}

// CustomVisionError represents the CustomVisionError schema from the OpenAPI specification
type CustomVisionError struct {
	Code string `json:"code"` // The error code.
	Message string `json:"message"` // A message explaining the error reported by the service.
}

// ImageUrlCreateBatch represents the ImageUrlCreateBatch schema from the OpenAPI specification
type ImageUrlCreateBatch struct {
	Images []ImageUrlCreateEntry `json:"images,omitempty"`
	Tagids []string `json:"tagIds,omitempty"`
}

// ImagePerformance represents the ImagePerformance schema from the OpenAPI specification
type ImagePerformance struct {
	Created string `json:"created,omitempty"`
	Id string `json:"id,omitempty"`
	Tags []ImageTag `json:"tags,omitempty"`
	Predictions []Prediction `json:"predictions,omitempty"`
	Height int `json:"height,omitempty"`
	Thumbnailuri string `json:"thumbnailUri,omitempty"`
	Width int `json:"width,omitempty"`
	Imageuri string `json:"imageUri,omitempty"`
	Regions []ImageRegion `json:"regions,omitempty"`
}

// ImageTagCreateBatch represents the ImageTagCreateBatch schema from the OpenAPI specification
type ImageTagCreateBatch struct {
	Tags []ImageTagCreateEntry `json:"tags,omitempty"` // Image Tag entries to include in this batch.
}

// IterationPerformance represents the IterationPerformance schema from the OpenAPI specification
type IterationPerformance struct {
	Averageprecision float32 `json:"averagePrecision,omitempty"` // Gets the average precision when applicable.
	Pertagperformance []TagPerformance `json:"perTagPerformance,omitempty"` // Gets the per-tag performance details for this iteration.
	Precision float32 `json:"precision,omitempty"` // Gets the precision.
	Precisionstddeviation float32 `json:"precisionStdDeviation,omitempty"` // Gets the standard deviation for the precision.
	Recall float32 `json:"recall,omitempty"` // Gets the recall.
	Recallstddeviation float32 `json:"recallStdDeviation,omitempty"` // Gets the standard deviation for the recall.
}

// Domain represents the Domain schema from the OpenAPI specification
type Domain struct {
	Exportable bool `json:"exportable,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	TypeField string `json:"type,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
}

// ImageFileCreateEntry represents the ImageFileCreateEntry schema from the OpenAPI specification
type ImageFileCreateEntry struct {
	Name string `json:"name,omitempty"`
	Regions []Region `json:"regions,omitempty"`
	Tagids []string `json:"tagIds,omitempty"`
	Contents string `json:"contents,omitempty"`
}

// ImageCreateSummary represents the ImageCreateSummary schema from the OpenAPI specification
type ImageCreateSummary struct {
	Images []ImageCreateResult `json:"images,omitempty"` // List of the image creation results.
	Isbatchsuccessful bool `json:"isBatchSuccessful,omitempty"` // True if all of the images in the batch were created successfully, otherwise false.
}

// ImagePrediction represents the ImagePrediction schema from the OpenAPI specification
type ImagePrediction struct {
	Id string `json:"id,omitempty"` // Prediction Id.
	Iteration string `json:"iteration,omitempty"` // Iteration Id.
	Predictions []Prediction `json:"predictions,omitempty"` // List of predictions.
	Project string `json:"project,omitempty"` // Project Id.
	Created string `json:"created,omitempty"` // Date this prediction was created.
}

// ImageCreateResult represents the ImageCreateResult schema from the OpenAPI specification
type ImageCreateResult struct {
	Image Image `json:"image,omitempty"` // Image model to be sent as JSON.
	Sourceurl string `json:"sourceUrl,omitempty"` // Source URL of the image.
	Status string `json:"status,omitempty"` // Status of the image creation.
}

// ImageIdCreateEntry represents the ImageIdCreateEntry schema from the OpenAPI specification
type ImageIdCreateEntry struct {
	Id string `json:"id,omitempty"` // Id of the image.
	Regions []Region `json:"regions,omitempty"`
	Tagids []string `json:"tagIds,omitempty"`
}

// ImageTag represents the ImageTag schema from the OpenAPI specification
type ImageTag struct {
	Tagid string `json:"tagId,omitempty"`
	Tagname string `json:"tagName,omitempty"`
	Created string `json:"created,omitempty"`
}

// RegionProposal represents the RegionProposal schema from the OpenAPI specification
type RegionProposal struct {
	Confidence float32 `json:"confidence,omitempty"`
	Boundingbox BoundingBox `json:"boundingBox,omitempty"` // Bounding box that defines a region of an image.
}

// ImageRegionProposal represents the ImageRegionProposal schema from the OpenAPI specification
type ImageRegionProposal struct {
	Imageid string `json:"imageId,omitempty"`
	Projectid string `json:"projectId,omitempty"`
	Proposals []RegionProposal `json:"proposals,omitempty"`
}

// Prediction represents the Prediction schema from the OpenAPI specification
type Prediction struct {
	Tagname string `json:"tagName,omitempty"` // Name of the predicted tag.
	Boundingbox BoundingBox `json:"boundingBox,omitempty"` // Bounding box that defines a region of an image.
	Probability float32 `json:"probability,omitempty"` // Probability of the tag.
	Tagid string `json:"tagId,omitempty"` // Id of the predicted tag.
}

// ImageRegionCreateBatch represents the ImageRegionCreateBatch schema from the OpenAPI specification
type ImageRegionCreateBatch struct {
	Regions []ImageRegionCreateEntry `json:"regions,omitempty"`
}

// PredictionQueryResult represents the PredictionQueryResult schema from the OpenAPI specification
type PredictionQueryResult struct {
	Token PredictionQueryToken `json:"token,omitempty"`
	Results []StoredImagePrediction `json:"results,omitempty"`
}

// TagPerformance represents the TagPerformance schema from the OpenAPI specification
type TagPerformance struct {
	Precisionstddeviation float32 `json:"precisionStdDeviation,omitempty"` // Gets the standard deviation for the precision.
	Recall float32 `json:"recall,omitempty"` // Gets the recall.
	Recallstddeviation float32 `json:"recallStdDeviation,omitempty"` // Gets the standard deviation for the recall.
	Averageprecision float32 `json:"averagePrecision,omitempty"` // Gets the average precision when applicable.
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Precision float32 `json:"precision,omitempty"` // Gets the precision.
}

// ProjectSettings represents the ProjectSettings schema from the OpenAPI specification
type ProjectSettings struct {
	Classificationtype string `json:"classificationType,omitempty"` // Gets or sets the classification type of the project.
	Domainid string `json:"domainId,omitempty"` // Gets or sets the id of the Domain to use with this project.
	Targetexportplatforms []string `json:"targetExportPlatforms,omitempty"` // A list of ExportPlatform that the trained model should be able to support.
}

// Export represents the Export schema from the OpenAPI specification
type Export struct {
	Downloaduri string `json:"downloadUri,omitempty"` // URI used to download the model.
	Flavor string `json:"flavor,omitempty"` // Flavor of the export.
	Newerversionavailable bool `json:"newerVersionAvailable,omitempty"` // Indicates an updated version of the export package is available and should be re-exported for the latest changes.
	Platform string `json:"platform,omitempty"` // Platform of the export.
	Status string `json:"status,omitempty"` // Status of the export.
}

// ImageIdCreateBatch represents the ImageIdCreateBatch schema from the OpenAPI specification
type ImageIdCreateBatch struct {
	Images []ImageIdCreateEntry `json:"images,omitempty"`
	Tagids []string `json:"tagIds,omitempty"`
}

// ImageRegion represents the ImageRegion schema from the OpenAPI specification
type ImageRegion struct {
	Created string `json:"created,omitempty"`
	Height float32 `json:"height"` // Height.
	Left float32 `json:"left"` // Coordinate of the left boundary.
	Regionid string `json:"regionId,omitempty"`
	Tagid string `json:"tagId"` // Id of the tag associated with this region.
	Tagname string `json:"tagName,omitempty"`
	Top float32 `json:"top"` // Coordinate of the top boundary.
	Width float32 `json:"width"` // Width.
}

// Region represents the Region schema from the OpenAPI specification
type Region struct {
	Width float32 `json:"width"` // Width.
	Height float32 `json:"height"` // Height.
	Left float32 `json:"left"` // Coordinate of the left boundary.
	Tagid string `json:"tagId"` // Id of the tag associated with this region.
	Top float32 `json:"top"` // Coordinate of the top boundary.
}

// ImageTagCreateEntry represents the ImageTagCreateEntry schema from the OpenAPI specification
type ImageTagCreateEntry struct {
	Imageid string `json:"imageId,omitempty"` // Id of the image.
	Tagid string `json:"tagId,omitempty"` // Id of the tag.
}

// ImageRegionCreateSummary represents the ImageRegionCreateSummary schema from the OpenAPI specification
type ImageRegionCreateSummary struct {
	Exceeded []ImageRegionCreateEntry `json:"exceeded,omitempty"`
	Created []ImageRegionCreateResult `json:"created,omitempty"`
	Duplicated []ImageRegionCreateEntry `json:"duplicated,omitempty"`
}

// ImageTagCreateSummary represents the ImageTagCreateSummary schema from the OpenAPI specification
type ImageTagCreateSummary struct {
	Exceeded []ImageTagCreateEntry `json:"exceeded,omitempty"`
	Created []ImageTagCreateEntry `json:"created,omitempty"`
	Duplicated []ImageTagCreateEntry `json:"duplicated,omitempty"`
}

// StoredImagePrediction represents the StoredImagePrediction schema from the OpenAPI specification
type StoredImagePrediction struct {
	Originalimageuri string `json:"originalImageUri,omitempty"` // The URI to the original prediction image.
	Resizedimageuri string `json:"resizedImageUri,omitempty"` // The URI to the (resized) prediction image.
	Project string `json:"project,omitempty"` // Project Id.
	Created string `json:"created,omitempty"` // Date this prediction was created.
	Domain string `json:"domain,omitempty"` // Domain used for the prediction.
	Predictions []Prediction `json:"predictions,omitempty"` // List of predictions.
	Thumbnailuri string `json:"thumbnailUri,omitempty"` // The URI to the thumbnail of the original prediction image.
	Id string `json:"id,omitempty"` // Prediction Id.
	Iteration string `json:"iteration,omitempty"` // Iteration Id.
}

// Image represents the Image schema from the OpenAPI specification
type Image struct {
	Tags []ImageTag `json:"tags,omitempty"` // Tags associated with this image.
	Width int `json:"width,omitempty"` // Width of the image.
	Created string `json:"created,omitempty"` // Date the image was created.
	Height int `json:"height,omitempty"` // Height of the image.
	Id string `json:"id,omitempty"` // Id of the image.
	Regions []ImageRegion `json:"regions,omitempty"` // Regions associated with this image.
	Resizedimageuri string `json:"resizedImageUri,omitempty"` // The URI to the (resized) image used for training.
	Thumbnailuri string `json:"thumbnailUri,omitempty"` // The URI to the thumbnail of the original image.
	Originalimageuri string `json:"originalImageUri,omitempty"` // The URI to the original uploaded image.
}

// ImageUrlCreateEntry represents the ImageUrlCreateEntry schema from the OpenAPI specification
type ImageUrlCreateEntry struct {
	Url string `json:"url"` // Url of the image.
	Regions []Region `json:"regions,omitempty"`
	Tagids []string `json:"tagIds,omitempty"`
}

// Iteration represents the Iteration schema from the OpenAPI specification
type Iteration struct {
	Domainid string `json:"domainId,omitempty"` // Get or sets a guid of the domain the iteration has been trained on.
	Exportable bool `json:"exportable,omitempty"` // Whether the iteration can be exported to another format for download.
	Projectid string `json:"projectId,omitempty"` // Gets the project id of the iteration.
	Status string `json:"status,omitempty"` // Gets the current iteration status.
	Trainedat string `json:"trainedAt,omitempty"` // Gets the time this iteration was last modified.
	Trainingtype string `json:"trainingType,omitempty"` // Gets the training type of the iteration.
	Id string `json:"id,omitempty"` // Gets the id of the iteration.
	Created string `json:"created,omitempty"` // Gets the time this iteration was completed.
	Exportableto []string `json:"exportableTo,omitempty"` // A set of platforms this iteration can export to.
	Originalpublishresourceid string `json:"originalPublishResourceId,omitempty"` // Resource Provider Id this iteration was originally published to.
	Name string `json:"name"` // Gets or sets the name of the iteration.
	Classificationtype string `json:"classificationType,omitempty"` // Gets the classification type of the project.
	Reservedbudgetinhours int `json:"reservedBudgetInHours,omitempty"` // Gets the reserved advanced training budget for the iteration.
	Publishname string `json:"publishName,omitempty"` // Name of the published model.
	Lastmodified string `json:"lastModified,omitempty"` // Gets the time this iteration was last modified.
}
