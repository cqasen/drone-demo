package entity

import "time"

const (
	TableCategory = "zbp_category"
	TableComment  = "zbp_comment"
	TableConfig   = "zbp_config"
	TableMember   = "zbp_member"
	TableModule   = "zbp_module"
	TablePost     = "zbp_post"
	TableTag      = "zbp_tag"
	TableUpload   = "zbp_upload"
	TableRouters  = "routers"
)

// ZbpCategory ...
type ZbpCategory struct {
	CateID          int32  `gorm:"primary_key;AUTO_INCREMENT;column:cate_ID" json:"cate_id"`
	CateName        string `gorm:"column:cate_Name" json:"cate_name"`
	CateOrder       int32  `gorm:"column:cate_Order" json:"cate_order"`
	CateType        int32  `gorm:"column:cate_Type" json:"cate_type"`
	CateCount       int32  `gorm:"column:cate_Count" json:"cate_count"`
	CateAlias       string `gorm:"column:cate_Alias" json:"cate_alias"`
	CateIntro       string `gorm:"column:cate_Intro" json:"cate_intro"`
	CateRootID      int32  `gorm:"column:cate_RootID" json:"cate_root_id"`
	CateParentID    int32  `gorm:"column:cate_ParentID" json:"cate_parent_id"`
	CateTemplate    string `gorm:"column:cate_Template" json:"cate_template"`
	CateLogTemplate string `gorm:"column:cate_LogTemplate" json:"cate_log_template"`
	CateMeta        string `gorm:"column:cate_Meta" json:"cate_meta"`
}

func (ZbpCategory) TableName() string {
	return TableCategory
}

// ZbpComment ...
type ZbpComment struct {
	CommID         int32  `gorm:"primary_key;AUTO_INCREMENT;column:comm_ID" json:"comm_id"`
	CommLogID      int32  `gorm:"column:comm_LogID" json:"comm_log_id"`
	CommIsChecking int8   `gorm:"column:comm_IsChecking" json:"comm_is_checking"`
	CommRootID     int32  `gorm:"column:comm_RootID" json:"comm_root_id"`
	CommParentID   int32  `gorm:"column:comm_ParentID" json:"comm_parent_id"`
	CommAuthorID   int32  `gorm:"column:comm_AuthorID" json:"comm_author_id"`
	CommName       string `gorm:"column:comm_Name" json:"comm_name"`
	CommEmail      string `gorm:"column:comm_Email" json:"comm_email"`
	CommHomePage   string `gorm:"column:comm_HomePage" json:"comm_home_page"`
	CommContent    string `gorm:"column:comm_Content" json:"comm_content"`
	CommPostTime   int32  `gorm:"column:comm_PostTime" json:"comm_post_time"`
	CommIP         string `gorm:"column:comm_IP" json:"comm_ip"`
	CommAgent      string `gorm:"column:comm_Agent" json:"comm_agent"`
	CommMeta       string `gorm:"column:comm_Meta" json:"comm_meta"`
}

func (ZbpComment) TableName() string {
	return TableComment
}

// ZbpConfig ...
type ZbpConfig struct {
	ConfID    int32  `gorm:"primary_key;AUTO_INCREMENT;column:conf_ID" json:"conf_id"`
	ConfName  string `gorm:"column:conf_Name" json:"conf_name"`
	ConfValue string `gorm:"column:conf_Value" json:"conf_value"`
}

func (ZbpConfig) TableName() string {
	return TableConfig
}

// ZbpMember ...
type ZbpMember struct {
	MemID       int32  `gorm:"primary_key;AUTO_INCREMENT;column:mem_ID" json:"mem_id"`
	MemGUID     string `gorm:"column:mem_Guid" json:"mem_guid"`
	MemLevel    int8   `gorm:"column:mem_Level" json:"mem_level"`
	MemStatus   int8   `gorm:"column:mem_Status" json:"mem_status"`
	MemName     string `gorm:"column:mem_Name" json:"mem_name"`
	MemPassword string `gorm:"column:mem_Password" json:"mem_password"`
	MemEmail    string `gorm:"column:mem_Email" json:"mem_email"`
	MemHomePage string `gorm:"column:mem_HomePage" json:"mem_home_page"`
	MemIP       string `gorm:"column:mem_IP" json:"mem_ip"`
	MemPostTime int32  `gorm:"column:mem_PostTime" json:"mem_post_time"`
	MemAlias    string `gorm:"column:mem_Alias" json:"mem_alias"`
	MemIntro    string `gorm:"column:mem_Intro" json:"mem_intro"`
	MemArticles int32  `gorm:"column:mem_Articles" json:"mem_articles"`
	MemPages    int32  `gorm:"column:mem_Pages" json:"mem_pages"`
	MemComments int32  `gorm:"column:mem_Comments" json:"mem_comments"`
	MemUploads  int32  `gorm:"column:mem_Uploads" json:"mem_uploads"`
	MemTemplate string `gorm:"column:mem_Template" json:"mem_template"`
	MemMeta     string `gorm:"column:mem_Meta" json:"mem_meta"`
}

func (ZbpMember) TableName() string {
	return TableMember
}

// ZbpModule ...
type ZbpModule struct {
	ModID          int32  `gorm:"primary_key;AUTO_INCREMENT;column:mod_ID" json:"mod_id"`
	ModName        string `gorm:"column:mod_Name" json:"mod_name"`
	ModFileName    string `gorm:"column:mod_FileName" json:"mod_file_name"`
	ModContent     string `gorm:"column:mod_Content" json:"mod_content"`
	ModSidebarID   int32  `gorm:"column:mod_SidebarID" json:"mod_sidebar_id"`
	ModHtmlID      string `gorm:"column:mod_HtmlID" json:"mod_html_id"`
	ModType        string `gorm:"column:mod_Type" json:"mod_type"`
	ModMaxLi       int32  `gorm:"column:mod_MaxLi" json:"mod_max_li"`
	ModSource      string `gorm:"column:mod_Source" json:"mod_source"`
	ModIsHideTitle int8   `gorm:"column:mod_IsHideTitle" json:"mod_is_hide_title"`
	ModMeta        string `gorm:"column:mod_Meta" json:"mod_meta"`
}

func (ZbpModule) TableName() string {
	return TableModule
}

// ZbpPost ...
type ZbpPost struct {
	LogID       int32  `gorm:"primary_key;AUTO_INCREMENT;column:log_ID" json:"log_id"`
	LogCateID   int32  `gorm:"column:log_CateID" json:"log_cate_id"`
	LogAuthorID int32  `gorm:"column:log_AuthorID" json:"log_author_id"`
	LogTag      string `gorm:"column:log_Tag" json:"log_tag"`
	LogStatus   int8   `gorm:"column:log_Status" json:"log_status"`
	LogType     int32  `gorm:"column:log_Type" json:"log_type"`
	LogAlias    string `gorm:"column:log_Alias" json:"log_alias"`
	LogIsTop    int32  `gorm:"column:log_IsTop" json:"log_is_top"`
	LogIsLock   int8   `gorm:"column:log_IsLock" json:"log_is_lock"`
	LogTitle    string `gorm:"column:log_Title" json:"log_title"`
	LogIntro    string `gorm:"column:log_Intro" json:"log_intro"`
	LogContent  string `gorm:"column:log_Content" json:"log_content"`
	LogPostTime int32  `gorm:"column:log_PostTime" json:"log_post_time"`
	LogCommNums int32  `gorm:"column:log_CommNums" json:"log_comm_nums"`
	LogViewNums int32  `gorm:"column:log_ViewNums" json:"log_view_nums"`
	LogTemplate string `gorm:"column:log_Template" json:"log_template"`
	LogMeta     string `gorm:"column:log_Meta" json:"log_meta"`
}

func (ZbpPost) TableName() string {
	return TablePost
}

// ZbpTag ...
type ZbpTag struct {
	TagID       int32  `gorm:"primary_key;AUTO_INCREMENT;column:tag_ID" json:"tag_id"`
	TagName     string `gorm:"column:tag_Name" json:"tag_name"`
	TagOrder    int32  `gorm:"column:tag_Order" json:"tag_order"`
	TagType     int32  `gorm:"column:tag_Type" json:"tag_type"`
	TagCount    int32  `gorm:"column:tag_Count" json:"tag_count"`
	TagAlias    string `gorm:"column:tag_Alias" json:"tag_alias"`
	TagIntro    string `gorm:"column:tag_Intro" json:"tag_intro"`
	TagTemplate string `gorm:"column:tag_Template" json:"tag_template"`
	TagMeta     string `gorm:"column:tag_Meta" json:"tag_meta"`
}

func (ZbpTag) TableName() string {
	return TableTag
}

// ZbpUpload ...
type ZbpUpload struct {
	UlID         int32  `gorm:"primary_key;AUTO_INCREMENT;column:ul_ID" json:"ul_id"`
	UlAuthorID   int32  `gorm:"column:ul_AuthorID" json:"ul_author_id"`
	UlSize       int32  `gorm:"column:ul_Size" json:"ul_size"`
	UlName       string `gorm:"column:ul_Name" json:"ul_name"`
	UlSourceName string `gorm:"column:ul_SourceName" json:"ul_source_name"`
	UlMimeType   string `gorm:"column:ul_MimeType" json:"ul_mime_type"`
	UlPostTime   int32  `gorm:"column:ul_PostTime" json:"ul_post_time"`
	UlDownNums   int32  `gorm:"column:ul_DownNums" json:"ul_down_nums"`
	UlLogID      int32  `gorm:"column:ul_LogID" json:"ul_log_id"`
	UlIntro      string `gorm:"column:ul_Intro" json:"ul_intro"`
	UlMeta       string `gorm:"column:ul_Meta" json:"ul_meta"`
}

func (ZbpUpload) TableName() string {
	return TableUpload
}

type Routers struct {
	ID         uint32    `json:"id"`
	Path       string    `json:"path"`
	Method     string    `json:"method"`
	Del        uint16    `json:"del"`
	Createtime time.Time `json:"createtime"`
	Updatetime time.Time `json:"updatetime"`
}

