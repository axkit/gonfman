drop table config_data_types;
drop table config_sections;
drop table config_controls;
drop table config_params;


create table config_data_types (
	id 				VARCHAR(32) not null,
	constraint PK_CFGDT primary key (id)
);

insert into config_data_types (id) values
('int'),
('bool'),
('string'),
('float');

create table config_controls (
   id                   VARCHAR(128)         not null,
   validation_function        TEXT,
   failed_validation_response TEXT,
   constraint PK_CFGCTRL primary key (id)
);

insert into config_controls(id) values
('int-input'),
('checkbox'),
('string-input'),
('textarea'),
('float-input'),
('money'),
('location'),
('ip'),
('dir'),
('file'),
('date'),
('color'),
('duration'),
('year'),
('month');



create table config_sections (
   id                   VARCHAR(128)         not null,
   parent_id            VARCHAR(128)         null,
   position_order       INT4                 not null default 0,
   name                 TEXT         		 not null,
   constraint PK_CFGSEC primary key (id)
);

create table config_params (
   id                   VARCHAR(128)         not null,
   section_id           VARCHAR(128)         null,
   position_order       INT4                 not null default 0,
   name                 text         		 not null,
   data_type_id         VARCHAR(32)          not null,
   control_id           VARCHAR(32)          ,
   raw_value            TEXT         		 null,
   is_readonly          BOOL                 default false,
   is_nullable          BOOL 			     default false,
   is_sensitive 		BOOL 			     default false,
   updated_at 			timestamptz 		 ,
   updated_finger_print bigint,   
   constraint PK_CFGPARAM primary key (id),
   constraint FK_CFGPARAM_CFGDT  foreign key (data_type_id) references config_data_types(id),
   constraint FK_CFGPARAM_CFGCTRL foreign key (control_id) references config_controls(id),
   constraint FK_CFGPARAM_CFGSEC foreign key (section_id) references config_sections(id)
);

/*
insert into production.parameters(id, section_id, position_order, name, type_id, control_id, raw_value, is_readonly) values
('database.timezone', )
('general', 					null,     	true, 	10, 'General', 						null, 		null, 		true, 'management-console'),
	('is_demo', 				'general', 	false, 	10, 'Is deployed in demo mode', 	'bool', 	'true', 	true, 'management-console'),
('frontend', 					null,     	true, 	20, 'Frontend', 					null, 		null, 		true, 'management-console'),
	('sign_out_forward_page', 	'frontend', false, 	10, 'Arrival page after sign out', 	'string', 	'/login', 	false, 'management-console'),
	('default_language', 		'frontend', false, 	20, 'Default language', 			'string', 	'ru', 		true, 'management-console'),
('mailer', 						null,     	true, 	30, 'Emailer', 						null, 		null, 		true, 'management-console'),
	('contact_email', 			'mailer',   false, 	10, 'Contact email', 				'string', 	'mail@mail.com', 	false, 'management-console'),
	('contact_subject', 		'mailer',   false, 	20, 'Email subject', 				'string', 	'', 		false, 'management-console'),
('backend', 					null,     	true, 	40, 'Database', 					 null,    	null, 		true, 'management-console'),
	('database', 			'backend',      true, 	10, 'Backend', 		null, 	null, 		null, 'management-console'),
	('database.max_open_conns', 			'database', false, 	10, 'Max open connections', 		'int32', 	'100', 		false, 'management-console'),
	('database.max_idle_conns', 			'database', false, 	20, 'Max idle connections', 		'int32', 	'3', 		false, 'management-console');
	('database.timezone', 			        'database', false, 	20, 'Max idle connections', 		'int32', 	'3', 		false, 'management-console');
*/

