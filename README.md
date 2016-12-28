# reminds
create table an_data(
	id int primary key not null auto_increment,
	title varchar(50) default "",
	content varchar(250) default "",
	create_time varchar(30) default "",
	last_time varchar(30) default ""
);

db:remind