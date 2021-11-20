create table if not exists public.accounts
(
    id                      serial primary key,
    "name"                  citext                             not null,
    -- Remember that the allowed email length is 254
    -- https://www.rfc-editor.org/errata_search.php?rfc=3696&eid=1690
    "email"                 citext
        constraint email_chk check (char_length(email) <= 254) not null,
    password_hash           text                               not null,
    password_salt           text                               null,
    password_hash_algorithm text                               not null
);