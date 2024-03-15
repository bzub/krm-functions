variable "REGISTRY" {default = "ghcr.io"}
variable "USERNAME" {default = "bzub"}
variable "REPO_NAME" {default = "krm-functions"}

group "default" {
  targets = [
    "remove-creationtimestamp",
  ]
}

target "remove-creationtimestamp" {
	# target = "remove-creationtimestamp"
	# dockerfile = "./remove-creationtimestamp/Dockerfile"
	context = "./remove-creationtimestamp"
	tags = [
		"${REGISTRY}/${USERNAME}/${REPO_NAME}/remove-creationtimestamp:v0.0.1"
	]
}
