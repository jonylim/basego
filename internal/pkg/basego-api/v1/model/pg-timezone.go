package model

// PgTimeZone contains a time zone's information.
type PgTimeZone struct {
	RedisNil  bool   `json:"-"         redis:"redisNil"`
	Name      string `json:"name"      redis:"name"`
	Abbrev    string `json:"abbrev"    redis:"abbrev"`
	UTCOffset string `json:"utcOffset" redis:"utcOffset"`
	IsDST     bool   `json:"isDST"     redis:"isDST"`
}
