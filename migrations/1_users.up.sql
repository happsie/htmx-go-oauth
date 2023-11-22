CREATE TABLE IF NOT EXISTS users
(
    id varchar
(
    255
) PRIMARY KEY NOT NULL,
    profile_image_url varchar
(
    255
) NOT NULL,
    login varchar
(
    255
) NOT NULL,
    display_name varchar
(
    255
) NOT NULL,
    email varchar
(
    255
) NOT NULL
    );