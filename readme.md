# REST API Example

REST API example using net/http, httprouter, and sql/go

```sql
create table if not exists movies (
	id int auto_increment primary KEY,
	title VARCHAR(255) not null,
	rating int not null
);
```