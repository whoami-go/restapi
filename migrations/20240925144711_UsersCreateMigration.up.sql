    CREATE TABLE users (
       id bigserial not null,
        login varchar not null,
        password varchar not null
    );

    CREATE TABLE articles (
      id bigserial not null primary key,
      title varchar not null unique,
      author varchar not null,
      content varchar not null

);
