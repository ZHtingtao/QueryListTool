CREATE TABLE `QUERY_CONF` (
  `key` varchar(100) NOT NULL COMMENT '唯一键，列表查询使用',
  `sql` varchar(255) NOT NULL COMMENT '查询数据使用的SQL语句',
  PRIMARY KEY (`key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='数据查询';

CREATE TABLE `COL_CONF` (
  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `key` varchar(100) NOT NULL COMMENT '列表唯一KEY',
  `name` varchar(100) NOT NULL COMMENT '查询表字段名',
  `display` int(2) NOT NULL DEFAULT '1' COMMENT '是否显示',
  `displayName` varchar(50) DEFAULT NULL COMMENT '显示名称',
  `condition` int(2) NOT NULL DEFAULT '1' COMMENT '是否作为条件查询',
  PRIMARY KEY (`id`,`key`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 COMMENT='数据显示列配置表';
