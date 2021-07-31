package setdata_common

type AclResource string

var (
	ResourceUsers AclResource = "users"
	ResourceAcl   AclResource = "acl"
	ResourcePost  AclResource = "post"
)

var (
	resources           = []AclResource{ResourceAcl, ResourceUsers, ResourcePost}
	stringToAclResource = map[string]AclResource{
		"users": ResourceUsers,
		"acl":   ResourceAcl,
		"Post":  ResourcePost,
	}
	aclResourceToString = map[AclResource]string{
		ResourceUsers: "users",
		ResourcePost:  "post",
		ResourceAcl:   "acl",
	}
)

func GetAclResourceString(resource AclResource) (string, error) {
	value, ok := aclResourceToString[resource]
	if !ok {
		return "", ErrCurrentAclResourceNotExist
	}
	return value, nil
}

func GetAclResource(resource string) (AclResource, error) {
	value, ok := stringToAclResource[resource]
	if !ok {
		return "", ErrCurrentAclResourceNotExist
	}
	return value, nil
}
