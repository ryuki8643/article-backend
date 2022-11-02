create table if not exists articles
(
    article_id serial,
    title text default 'no_title',
    article_content text default 'no_content',
    auther text default 'no_auther',
    primary key (article_id)
    );
insert into articles  (title,article_content) values ('no_title','no_content');

