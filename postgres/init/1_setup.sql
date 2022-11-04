drop table if exists codes;
drop table if exists steps;
drop table if exists articles;
create table if not exists articles
(
    article_id int,
    title text default 'no_title',
    author text default 'no_author',
    likes int default 0,
    primary key (article_id)
    );

create table if not exists steps
(
    step_primary_key int,
    article_id int,
    step_id int,
    article_content text,
    foreign key (article_id) references articles(article_id),
    primary key (step_primary_key)
    );

create table if not exists codes
(
    step_primary_key int,
    code_id int,
    code_file_name text default 'no_file',
    code_content text default 'no_content',

    foreign key (step_primary_key) references steps(step_primary_key),
    primary key (step_primary_key,code_id)
    );
insert into articles values (0,'no_title','no_author',0);
insert into steps values (0,0,0,'content'),(1,0,1,'content');
insert into codes values (0,0,'no_file','no_code'),(0,1,'no_file','no_code'),(1,0,'no_file','no_code'),(1,1,'no_file','exit');
