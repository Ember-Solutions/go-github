package github

import "time"

type PackageV2 struct {
	Id             *int64            `json:"id"`
	Name           *string           `json:"name"`
	Namespace      *string           `json:"namespace"`
	Description    *string           `json:"description"`
	Ecosystem      *string           `json:"ecosystem"`
	HtmlUrl        *string           `json:"html_url"`
	CreatedAt      *time.Time        `json:"created_at"`
	UpdatedAt      *time.Time        `json:"updated_at"`
	PackageVersion *PackageVersionV2 `json:"package_version"`
}

func (p PackageV2) String() string {
	return Stringify(p)
}

type PackageVersionV2 struct {
	Id          *int64                `json:"id"`
	Name        *string               `json:"name"`
	Description *string               `json:"description"`
	BlobStore   *string               `json:"blob_store"`
	HtmlUrl     *string               `json:"html_url"`
	CreatedAt   *time.Time            `json:"created_at"`
	UpdatedAt   *time.Time            `json:"updated_at"`
	NPMMetadata *PackageV2NPMMetadata `json:"npm_metadata"`
}

func (p PackageVersionV2) String() string {
	return Stringify(p)
}

type PackageV2NPMMetadata struct {
	Name    *string `json:"name"`
	Version *string `json:"version"`
	NpmUser *string `json:"npm_user"`
	Author  *struct {
		Name  *string `json:"name"`
		Email *string `json:"email"`
		Url   *string `json:"url"`
	} `json:"author"`
	Bugs *struct {
		Url *string `json:"url"`
	} `json:"bugs"`
	Dependencies         map[string]string `json:"dependencies"`
	DevDependencies      map[string]string `json:"dev_dependencies"`
	PeerDependencies     map[string]string `json:"peer_dependencies"`
	OptionalDependencies map[string]string `json:"optional_dependencies"`
	Description          *string           `json:"description"`
	Dist                 *struct {
		Integrity *string `json:"integrity"`
		Shasum    *string `json:"shasum"`
		Tarball   *string `json:"tarball"`
	} `json:"dist"`
	GitHead    *string `json:"git_head"`
	Homepage   *string `json:"homepage"`
	License    *string `json:"license"`
	Main       *string `json:"main"`
	Repository *struct {
		Type      *string `json:"type"`
		Url       *string `json:"url"`
		Directory *string `json:"directory"`
	} `json:"repository"`
	Scripts             map[string]string `json:"scripts"`
	Id                  *string           `json:"id"`
	NodeVersion         *string           `json:"node_version"`
	NpmVersion          *string           `json:"npm_version"`
	HasShrinkwrap       *bool             `json:"has_shrinkwrap"`
	Maintainers         []interface{}     `json:"maintainers"`
	Contributors        []interface{}     `json:"contributors"`
	Engines             map[string]string `json:"engines"`
	Keywords            []interface{}     `json:"keywords"`
	Files               []string          `json:"files"`
	Readme              *string           `json:"readme"`
	InstallationCommand *string           `json:"installation_command"`
	ReleaseId           *int              `json:"release_id"`
	CommitOid           *string           `json:"commit_oid"`
	PublishedViaActions *bool             `json:"published_via_actions"`
	DeletedById         *int              `json:"deleted_by_id"`
}

func (p PackageV2NPMMetadata) String() string {
	return Stringify(p)
}
