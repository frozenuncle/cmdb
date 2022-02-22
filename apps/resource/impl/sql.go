package impl

const (
	SQLInsertResource = `INSERT INTO resource (
		id,vendor,region,zone,create_at,expire_at,category,type,
		name,description,status,update_at,sync_at,sync_accout,public_ip,
		private_ip,pay_type,describe_hash,resource_hash,secret_id
	) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`

	SQLUpdateResource = `UPDATE resource SET 
		expire_at=?,category=?,type=?,name=?,description=?,
		status=?,update_at=?,sync_at=?,sync_accout=?,
		public_ip=?,private_ip=?,pay_type=?,describe_hash=?,resource_hash=?,
		secret_id=? 
	WHERE id = ?`

	SQLQueryResource     = `SELECT * FROM resource`
	SQLDeleteResource    = `DELETE FROM resource WHERE id = ?;`
	SQLDeleteResourceTag = `DELETE FROM tag WHERE resource_id = ?;`

	SQLInsertTag = `INSERT INTO tag (
		key,value,describe,resource_id
	) VALUES (?,?,?,?) ON DUPLICATE KEY UPDATE key=key;`
)
