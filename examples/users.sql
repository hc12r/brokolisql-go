CREATE TABLE "geos" (
  "id" INTEGER PRIMARY KEY,
  "lat" TEXT,
  "lng" TEXT
);

CREATE TABLE "addresses" (
  "id" INTEGER PRIMARY KEY,
  "geo_id" INTEGER,
  "street" TEXT,
  "suite" TEXT,
  "zipcode" TEXT,
  "city" TEXT,
  FOREIGN KEY ("geo_id") REFERENCES "geos" ("id") ON DELETE CASCADE
);

CREATE TABLE "companies" (
  "id" INTEGER PRIMARY KEY,
  "bs" TEXT,
  "catchPhrase" TEXT,
  "name" TEXT
);

CREATE TABLE "userses" (
  "id" INTEGER PRIMARY KEY,
  "name" TEXT,
  "username" TEXT,
  "email" TEXT,
  "address_id" INTEGER,
  "phone" TEXT,
  "website" TEXT,
  "company_id" INTEGER,
  FOREIGN KEY ("address_id") REFERENCES "addresses" ("id") ON DELETE CASCADE,
  FOREIGN KEY ("company_id") REFERENCES "companies" ("id") ON DELETE CASCADE
);

INSERT INTO "geos" ("id", "lat", "lng") VALUES
(31, '-37.3159', '81.1496'),
(32, '-43.9509', '-34.4618'),
(33, '-68.6102', '-47.0653'),
(34, '29.4572', '-164.2990'),
(35, '-31.8129', '62.5342'),
(36, '-71.4197', '71.7478'),
(37, '24.8918', '21.8984'),
(38, '-14.3990', '-120.7677'),
(39, '24.6463', '-168.8889'),
(40, '-38.2386', '57.2232');


INSERT INTO "addresses" ("id", "geo_id", "street", "suite", "zipcode", "city") VALUES
(11, 31, 'Kulas Light', 'Apt. 556', '92998-3874', 'Gwenborough'),
(12, 32, 'Victor Plains', 'Suite 879', '90566-7771', 'Wisokyburgh'),
(13, 33, 'Douglas Extension', 'Suite 847', '59590-4157', 'McKenziehaven'),
(14, 34, 'Hoeger Mall', 'Apt. 692', '53919-4257', 'South Elvis'),
(15, 35, 'Skiles Walks', 'Suite 351', '33263', 'Roscoeview'),
(16, 36, 'Norberto Crossing', 'Apt. 950', '23505-1337', 'South Christy'),
(17, 37, 'Rex Trail', 'Suite 280', '58804-1099', 'Howemouth'),
(18, 38, 'Ellsworth Summit', 'Suite 729', '45169', 'Aliyaview'),
(19, 39, 'Dayna Park', 'Suite 449', '76495-3109', 'Bartholomebury'),
(20, 40, 'Kattie Turnpike', 'Suite 198', '31428-2261', 'Lebsackbury');


INSERT INTO "companies" ("id", "bs", "catchPhrase", "name") VALUES
(21, 'harness real-time e-markets', 'Multi-layered client-server neural-net', 'Romaguera-Crona'),
(22, 'synergize scalable supply-chains', 'Proactive didactic contingency', 'Deckow-Crist'),
(23, 'e-enable strategic applications', 'Face to face bifurcated interface', 'Romaguera-Jacobson'),
(24, 'transition cutting-edge web services', 'Multi-tiered zero tolerance productivity', 'Robel-Corkery'),
(25, 'revolutionize end-to-end systems', 'User-centric fault-tolerant solution', 'Keebler LLC'),
(26, 'e-enable innovative applications', 'Synchronised bottom-line interface', 'Considine-Lockman'),
(27, 'generate enterprise e-tailers', 'Configurable multimedia task-force', 'Johns Group'),
(28, 'e-enable extensible e-tailers', 'Implemented secondary concept', 'Abernathy Group'),
(29, 'aggregate real-time technologies', 'Switchable contextually-based project', 'Yost and Sons'),
(30, 'target end-to-end models', 'Centralized empowering task-force', 'Hoeger LLC');


INSERT INTO "userses" ("id", "name", "username", "email", "address_id", "phone", "website", "company_id") VALUES
(1, 'Leanne Graham', 'Bret', 'Sincere@april.biz', 11, '1-770-736-8031 x56442', 'hildegard.org', 21),
(2, 'Ervin Howell', 'Antonette', 'Shanna@melissa.tv', 12, '010-692-6593 x09125', 'anastasia.net', 22),
(3, 'Clementine Bauch', 'Samantha', 'Nathan@yesenia.net', 13, '1-463-123-4447', 'ramiro.info', 23),
(4, 'Patricia Lebsack', 'Karianne', 'Julianne.OConner@kory.org', 14, '493-170-9623 x156', 'kale.biz', 24),
(5, 'Chelsey Dietrich', 'Kamren', 'Lucio_Hettinger@annie.ca', 15, '(254)954-1289', 'demarco.info', 25),
(6, 'Mrs. Dennis Schulist', 'Leopoldo_Corkery', 'Karley_Dach@jasper.info', 16, '1-477-935-8478 x6430', 'ola.org', 26),
(7, 'Kurtis Weissnat', 'Elwyn.Skiles', 'Telly.Hoeger@billy.biz', 17, '210.067.6132', 'elvis.io', 27),
(8, 'Nicholas Runolfsdottir V', 'Maxime_Nienow', 'Sherwood@rosamond.me', 18, '586.493.6943 x140', 'jacynthe.com', 28),
(9, 'Glenna Reichert', 'Delphine', 'Chaim_McDermott@dana.io', 19, '(775)976-6794 x41206', 'conrad.com', 29),
(10, 'Clementina DuBuque', 'Moriah.Stanton', 'Rey.Padberg@karina.biz', 20, '024-648-3804', 'ambrose.net', 30);


