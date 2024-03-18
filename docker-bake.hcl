variable "REGISTRY" {default = "ghcr.io"}
variable "USERNAME" {default = "bzub"}
variable "REPO_NAME" {default = "krm-functions"}

group "default" {
  targets = [
    "remove-creationtimestamp",
    "remove-cabundle",
  ]
}

target "remove-creationtimestamp" {
	context = "./remove-creationtimestamp"
	tags = [
		"${REGISTRY}/${USERNAME}/${REPO_NAME}/remove-creationtimestamp:v0.0.1"
	]
}

target "remove-cabundle" {
	context = "./remove-cabundle"
	tags = [
		"${REGISTRY}/${USERNAME}/${REPO_NAME}/remove-cabundle:v0.0.1"
	]
}
