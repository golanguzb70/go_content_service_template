CREATE TYPE language_code AS ENUM (
  'uz',
  'ru',
  'en'
);

CREATE TYPE video_status AS ENUM (
  'original',
  'inprocess',
  'transcoded'
);

CREATE TYPE movie_type AS ENUM (
  'movie',
  'series'
);

CREATE TYPE payment_type AS ENUM (
  'free',
  'tvod',
  'svod'
);

CREATE TABLE positions (
  id uuid PRIMARY KEY,
  title_uz varchar NOT NULL,
  title_ru varchar NOT NULL,
  title_en varchar NOT NULL,
  active bool NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP WITHOUT TIME ZONE
);

CREATE TABLE staffs (
  id uuid PRIMARY KEY,
  photo varchar NOT NULL,
  position_id uuid NOT NULL,
  slug varchar UNIQUE NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP WITHOUT TIME ZONE
);

CREATE TABLE stuff_translates (
  id uuid PRIMARY KEY,
  lang language_code NOT NULL,
  first_name varchar NOT NULL,
  last_name varchar NOT NULL,
  biography text NOT NULL,
  staff_id uuid NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP WITHOUT TIME ZONE
);