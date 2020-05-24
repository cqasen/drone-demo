package entity

const (
	TABLE_CATEGORY = "zbp_category"
	TABLE_COMMENT  = "zbp_comment"
	TABLE_CONFIG   = "zbp_config"
	TABLE_MEMBER   = "zbp_member"
	TABLE_MODULE   = "zbp_module"
	TABLE_POST     = "zbp_post"
	TABLE_TAG      = "zbp_tag"
	TABLE_UPLOAD   = "zbp_upload"
)

// ZbpCategory ...
type ZbpCategory struct {
	CateID          int32  `gorm:"primary_key;AUTO_INCREMENT;column:cate_ID"`
	CateName        string `gorm:"column:cate_Name"`
	CateOrder       int32  `gorm:"column:cate_Order"`
	CateType        int32  `gorm:"column:cate_Type"`
	CateCount       int32  `gorm:"column:cate_Count"`
	CateAlias       string `gorm:"column:cate_Alias"`
	CateIntro       string `gorm:"column:cate_Intro"`
	CateRootID      int32  `gorm:"column:cate_RootID"`
	CateParentID    int32  `gorm:"column:cate_ParentID"`
	CateTemplate    string `gorm:"column:cate_Template"`
	CateLogTemplate string `gorm:"column:cate_LogTemplate"`
	CateMeta        string `gorm:"column:cate_Meta"`
}

func (ZbpCategory) TableName() string {
	return TABLE_CATEGORY
}

// ZbpComment ...
type ZbpComment struct {
	CommID         int32  `gorm:"primary_key;AUTO_INCREMENT;column:comm_ID"`
	CommLogID      int32  `gorm:"column:comm_LogID"`
	CommIsChecking int8   `gorm:"column:comm_IsChecking"`
	CommRootID     int32  `gorm:"column:comm_RootID"`
	CommParentID   int32  `gorm:"column:comm_ParentID"`
	CommAuthorID   int32  `gorm:"column:comm_AuthorID"`
	CommName       string `gorm:"column:comm_Name"`
	CommEmail      string `gorm:"column:comm_Email"`
	CommHomePage   string `gorm:"column:comm_HomePage"`
	CommContent    string `gorm:"column:comm_Content"`
	CommPostTime   int32  `gorm:"column:comm_PostTime"`
	CommIP         string `gorm:"column:comm_IP"`
	CommAgent      string `gorm:"column:comm_Agent"`
	CommMeta       string `gorm:"column:comm_Meta"`
}

func (ZbpComment) TableName() string {
	return TABLE_COMMENT
}

// ZbpConfig ...
type ZbpConfig struct {
	ConfID    int32  `gorm:"primary_key;AUTO_INCREMENT;column:conf_ID"`
	ConfName  string `gorm:"column:conf_Name"`
	ConfValue string `gorm:"column:conf_Value"`
}

func (ZbpConfig) TableName() string {
	return TABLE_CONFIG
}

// ZbpMember ...
type ZbpMember struct {
	MemID       int32  `gorm:"primary_key;AUTO_INCREMENT;column:mem_ID"`
	MemGUID     string `gorm:"column:mem_Guid"`
	MemLevel    int8   `gorm:"column:mem_Level"`
	MemStatus   int8   `gorm:"column:mem_Status"`
	MemName     string `gorm:"column:mem_Name"`
	MemPassword string `gorm:"column:mem_Password"`
	MemEmail    string `gorm:"column:mem_Email"`
	MemHomePage string `gorm:"column:mem_HomePage"`
	MemIP       string `gorm:"column:mem_IP"`
	MemPostTime int32  `gorm:"column:mem_PostTime"`
	MemAlias    string `gorm:"column:mem_Alias"`
	MemIntro    string `gorm:"column:mem_Intro"`
	MemArticles int32  `gorm:"column:mem_Articles"`
	MemPages    int32  `gorm:"column:mem_Pages"`
	MemComments int32  `gorm:"column:mem_Comments"`
	MemUploads  int32  `gorm:"column:mem_Uploads"`
	MemTemplate string `gorm:"column:mem_Template"`
	MemMeta     string `gorm:"column:mem_Meta"`
}

func (ZbpMember) TableName() string {
	return TABLE_MEMBER
}

// ZbpModule ...
type ZbpModule struct {
	ModID          int32  `gorm:"primary_key;AUTO_INCREMENT;column:mod_ID"`
	ModName        string `gorm:"column:mod_Name"`
	ModFileName    string `gorm:"column:mod_FileName"`
	ModContent     string `gorm:"column:mod_Content"`
	ModSidebarID   int32  `gorm:"column:mod_SidebarID"`
	ModHtmlID      string `gorm:"column:mod_HtmlID"`
	ModType        string `gorm:"column:mod_Type"`
	ModMaxLi       int32  `gorm:"column:mod_MaxLi"`
	ModSource      string `gorm:"column:mod_Source"`
	ModIsHideTitle int8   `gorm:"column:mod_IsHideTitle"`
	ModMeta        string `gorm:"column:mod_Meta"`
}

func (ZbpModule) TableName() string {
	return TABLE_MODULE
}

// ZbpPost ...
type ZbpPost struct {
	LogID       int32  `gorm:"primary_key;AUTO_INCREMENT;column:log_ID"`
	LogCateID   int32  `gorm:"column:log_CateID"`
	LogAuthorID int32  `gorm:"column:log_AuthorID"`
	LogTag      string `gorm:"column:log_Tag"`
	LogStatus   int8   `gorm:"column:log_Status"`
	LogType     int32  `gorm:"column:log_Type"`
	LogAlias    string `gorm:"column:log_Alias"`
	LogIsTop    int32  `gorm:"column:log_IsTop"`
	LogIsLock   int8   `gorm:"column:log_IsLock"`
	LogTitle    string `gorm:"column:log_Title"`
	LogIntro    string `gorm:"column:log_Intro"`
	LogContent  string `gorm:"column:log_Content"`
	LogPostTime int32  `gorm:"column:log_PostTime"`
	LogCommNums int32  `gorm:"column:log_CommNums"`
	LogViewNums int32  `gorm:"column:log_ViewNums"`
	LogTemplate string `gorm:"column:log_Template"`
	LogMeta     string `gorm:"column:log_Meta"`
}

func (ZbpPost) TableName() string {
	return TABLE_POST
}

// ZbpTag ...
type ZbpTag struct {
	TagID       int32  `gorm:"primary_key;AUTO_INCREMENT;column:tag_ID"`
	TagName     string `gorm:"column:tag_Name"`
	TagOrder    int32  `gorm:"column:tag_Order"`
	TagType     int32  `gorm:"column:tag_Type"`
	TagCount    int32  `gorm:"column:tag_Count"`
	TagAlias    string `gorm:"column:tag_Alias"`
	TagIntro    string `gorm:"column:tag_Intro"`
	TagTemplate string `gorm:"column:tag_Template"`
	TagMeta     string `gorm:"column:tag_Meta"`
}

func (ZbpTag) TableName() string {
	return TABLE_TAG
}

// ZbpUpload ...
type ZbpUpload struct {
	UlID         int32  `gorm:"primary_key;AUTO_INCREMENT;column:ul_ID"`
	UlAuthorID   int32  `gorm:"column:ul_AuthorID"`
	UlSize       int32  `gorm:"column:ul_Size"`
	UlName       string `gorm:"column:ul_Name"`
	UlSourceName string `gorm:"column:ul_SourceName"`
	UlMimeType   string `gorm:"column:ul_MimeType"`
	UlPostTime   int32  `gorm:"column:ul_PostTime"`
	UlDownNums   int32  `gorm:"column:ul_DownNums"`
	UlLogID      int32  `gorm:"column:ul_LogID"`
	UlIntro      string `gorm:"column:ul_Intro"`
	UlMeta       string `gorm:"column:ul_Meta"`
}

func (ZbpUpload) TableName() string {
	return TABLE_UPLOAD
}
