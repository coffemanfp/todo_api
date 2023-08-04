create table if not exists account (
    id serial unique not null,
    name varchar not null,
    last_name varchar not null,
    nickname varchar unique not null,
    email varchar unique not null,
    password varchar not null,
    avatar_url varchar,
    created_at varchar,

    primary key (id)
);

create table if not exists list_category (
    id serial unique not null,
    name varchar not null,
    description varchar,
    icon_url varchar not null,
    created_by integer not null,

    primary key (id),
    foreign key (created_by) references account(id) ON DELETE CASCADE
);

create table if not exists list (
    id serial unique not null,
    title varchar not null,
    description varchar,
    background_picture_url varchar,
    created_at timestamp not null,
    created_by integer not null,

    primary key (id),
    foreign key (created_by) references account(id) ON DELETE CASCADE
);

create table if not exists list_category_union (
    id serial unique not null,
    list_id integer not null,
    category_id integer not null,
    created_by integer not null,

    primary key (id),
    foreign key (list_id) references list(id) ON DELETE CASCADE,
    foreign key (category_id) references list_category(id) ON DELETE CASCADE
);

create table if not exists task (
    id serial unique not null,
    created_by integer not null,
    list_id integer,
    title varchar not null,
    description varchar,
    is_important boolean,
    is_added_to_my_day boolean,
    reminder timestamp,
    due_date timestamp,
    repeat varchar,
    created_at timestamp not null,

    primary key (id),
    foreign key (list_id) references list(id) ON DELETE CASCADE,
    foreign key (created_by) references account(id) ON DELETE CASCADE
);

create table if not exists account_preferences (
    id serial unique not null,
    account_id integer not null,
    preferences jsonb not null,

    primary key (id),
    foreign key (account_id) references account(id) ON DELETE CASCADE
);