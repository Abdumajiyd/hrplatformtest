-- Foydalanuvchilar jadvali
CREATE TABLE users (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    phone_number VARCHAR NOT NULL,
    birthday TIMESTAMP,
    gender VARCHAR(1) CHECK (gender IN ('m', 'f')), -- ENUM o'rniga CHECK constraint
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

-- nega check constraint bo'ldi? 
-- PostgreSQL to'g'ridan-to'g'ri ENUM turlarini qo'llab-quvvatlamaydi.
-- Ya'ni, boshqa ba'zi ma'lumotlar bazalarida bo'lgani kabi masalan, MySQL, ENUM('qiymat1', 'qiymat2', ...)
-- ko'rinishida yangi ma'lumot turini yaratib bo'lmaydi. 
-- garchi CREATE TYPE ... AS ENUM sintaksisi yordamida ENUMga o'xshash tur yaratish mumkin bo'lsa-da, bu to'liq ENUM emas va ba'zi cheklovlarga ega.
-- (masalan, yangi qiymat qo'shish yoki olib tashlash qiyinroq).
-- ENUM turini o'zgartirish (yangi qiymat qo'shish yoki o'chirish) ancha murakkab jarayon bo'lishi mumkin. 
-- CHECK constraintlarni o'zgartirish esa oddiyroq va xavfsizroq.


-- Rezyumelar jadvali
CREATE TABLE resumes (
    id UUID PRIMARY KEY,
    position VARCHAR,
    experience INT,
    description TEXT,
    user_id UUID REFERENCES users(id)
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

-- Kompaniyalar jadvali
CREATE TABLE companies (
    id UUID PRIMARY KEY,
    name VARCHAR,
    location VARCHAR,
    workers INT
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

-- Rekruiterlar jadvali
CREATE TABLE recruiters (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    phone_number VARCHAR NOT NULL,
    birthday TIMESTAMP,
    gender VARCHAR(1) CHECK (gender IN ('m', 'f')), -- ENUM o'rniga CHECK constraint
    company_id UUID REFERENCES companies(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

-- Vakansiyalar jadvali
CREATE TABLE vacancies (
    id UUID PRIMARY KEY,
    name VARCHAR,
    position VARCHAR,
    min_exp INT,
    company_id UUID REFERENCES companies(id),
    description TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);

-- Intervyular jadvali
CREATE TABLE interviews (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    vacancy_id UUID REFERENCES vacancies(id),
    recruiter_id UUID REFERENCES recruiters(id),
    interview_date TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at BIGINT DEFAULT 0
);


-- ALTER TABLE resumes
-- ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT now();

-- ALTER TABLE resumes
-- ADD COLUMN updated_at TIMESTAMP NOT NULL DEFAULT now();

-- ALTER TABLE resumes
-- ADD COLUMN deleted_at BIGINT NOT NULL DEFAULT 0;

-- ALTER TABLE companies
-- ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT now();

-- ALTER TABLE companies
-- ADD COLUMN updated_at TIMESTAMP NOT NULL DEFAULT now();

-- ALTER TABLE companies
-- ADD COLUMN deleted_at BIGINT NOT NULL DEFAULT 0;
