CREATE TABLE "customers" (
  "ID" INTEGER,
  "GIVEN_NAME" TEXT,
  "SURNAME" TEXT,
  "EMAIL" TEXT,
  "COUNTRY" TEXT,
  "CITY" TEXT,
  "CONTACT_PHONE" TEXT,
  "SUBSCRIPTION_DATE" DATE,
  "FULL_NAME" TEXT
);

INSERT INTO "customers" ("ID", "GIVEN_NAME", "SURNAME", "EMAIL", "COUNTRY", "CITY", "CONTACT_PHONE", "SUBSCRIPTION_DATE", "FULL_NAME") VALUES
('3', 'Michael', 'Johnson', 'michael.j@example.com', 'Canada', 'Toronto', '+1-416-555-7890', '2023-01-10', 'Michael'' ''Johnson'),
('9', 'William', 'Thomas', 'william.t@example.com', 'Canada', 'Vancouver', '+1-604-555-1234', '2023-01-30', 'William'' ''Thomas'),
('6', 'Sarah', 'Taylor', 'sarah.taylor@example.com', 'France', 'Paris', '+33-1-2345-6789', '2023-01-22', 'Sarah'' ''Taylor'),
('5', 'David', 'Wilson', 'david.wilson@example.com', 'Germany', 'Berlin', '+49-30-1234-5678', '2023-02-28', 'David'' ''Wilson'),
('2', 'Jane', 'Smith', 'jane.smith@example.com', 'United Kingdom', 'London', '+44-20-1234-5678', '2023-02-20', 'Jane'' ''Smith'),
('7', 'Robert', 'Anderson', 'robert.a@example.com', 'United States', 'Chicago', '+1-312-555-6789', '2023-03-12', 'Robert'' ''Anderson'),
('1', 'John', 'Doe', 'john.doe@example.com', 'United States', 'New York', '+1-555-123-4567', '2023-01-15', 'John'' ''Doe');

