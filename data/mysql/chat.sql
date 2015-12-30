use chat;

#创建用户表
drop table if exists t_client;

create table t_client(
	clientId integer not null auto_increment,
	name varchar(256) not null,
	password varchar(256) not null,
	createTime timestamp not null default CURRENT_TIMESTAMP,
	modifyTime timestamp not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
	primary key( clientId )
)engine=innodb default charset=utf8mb4 auto_increment = 10001;

alter table t_client add index nameAndPasswordIndex(name, password);

#创建聊天信息表
drop table if exists t_message;

create table t_message(
	messageId integer not null auto_increment,
	text varchar(2048) not null,
	sendClientId integer not null,
	receiveClientId integer not null,
	type integer not null,
	createTime timestamp not null default CURRENT_TIMESTAMP,
	modifyTime timestamp not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
	primary key( messageId )
)engine=innodb default charset=utf8mb4 auto_increment = 10001;

alter table t_message add index receiveClientIdIndex(receiveClientId);
alter table t_message add index sendClientIndex(sendClientId);
alter table t_message add index receiveClientIdAndSendClientIdIndex(receiveClientId, sendClientId);

#创建刷新请求表
drop table if exists t_request;

create table t_request(
	requestId integer not null auto_increment,
	clientId integer not null,
	requestTime timestamp not null,
	createTime timestamp not null default CURRENT_TIMESTAMP,
	modifyTime timestamp not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
	primary key( requestId )
)engine=innodb default charset=utf8mb4 auto_increment = 10001;

alter table t_request add index clientIdIndex(clientId);

#初始数据
insert into t_client(clientId, name, password)values
(10001, 'jd', sha1('123')),
(10002, 'fish', sha1('123'));

select * from t_client;
