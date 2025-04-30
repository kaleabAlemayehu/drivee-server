-- +goose Up
-- +goose StatementBegin
INSERT INTO cars (
  id, owner_id, thumbnail_picture, description, make, model, year,
  license_plate, vin_number, transmission, fuel_type, mileage, location,
  price_per_hour, status, created_at, updated_at
) VALUES (
  '52bd5d9d-c2c3-474c-9507-e883f65862c1', '74e957d3-34b2-468a-95dd-fa816da9d5a3', 'http://localhost:9090/static/carphoto1.png', 'Mean care accept religious hotel evidence mean.', 'Toyota', 'Camry', '2020',
  'A45-OLD', 'fzV4116077UAbvDqE', 'manual', 'gas', 58595,
  ST_SetSRID(ST_MakePoint(-25.915308, 85.822688), 4326), 36.77, 'inactive',
  '2025-04-29 20:30:41', '2025-04-29 20:30:41'
);

INSERT INTO cars (
  id, owner_id, thumbnail_picture, description, make, model, year,
  license_plate, vin_number, transmission, fuel_type, mileage, location,
  price_per_hour, status, created_at, updated_at
) VALUES (
  '6289be13-c329-42f7-8499-c5f2afe637b4', '43fa3157-c671-4fd7-aed3-c165ab44fc98', 'http://localhost:9090/static/carphoto2.png', 'Yet attorney coach threat wind town store little method instead.', 'Honda', 'Civic', '2018',
  '2GK L05', 'iYi9871836DoQPZNO', 'automatic', 'gas', 108125,
  ST_SetSRID(ST_MakePoint(-121.664535, -22.0013315), 4326), 34.28, 'inactive',
  '2025-04-29 20:30:41', '2025-04-29 20:30:41'
);

INSERT INTO cars (
  id, owner_id, thumbnail_picture, description, make, model, year,
  license_plate, vin_number, transmission, fuel_type, mileage, location,
  price_per_hour, status, created_at, updated_at
) VALUES (
  'f93b89b4-e585-4f5a-a879-2d3930b046dc', 'a02da0df-2e4e-4c1b-8d42-29a627ab6ac3', 'http://localhost:9090/static/carphoto3.png', 'Develop hold house wall brother phone.', 'Ford', 'Mustang', '2015',
  '657N', 'RTp5740161idUUQzQ', 'automatic', 'gas', 127791,
  ST_SetSRID(ST_MakePoint(13.257346, 67.053208), 4326), 14.12, 'avaliable',
  '2025-04-29 20:30:41', '2025-04-29 20:30:41'
);

INSERT INTO cars (
  id, owner_id, thumbnail_picture, description, make, model, year,
  license_plate, vin_number, transmission, fuel_type, mileage, location,
  price_per_hour, status, created_at, updated_at
) VALUES (
  'b5808143-5a59-4017-bb5a-79d1fefe048d', '854a9efa-4cbf-4ffa-ab1f-fb419bcaf7f9', 'http://localhost:9090/static/carphoto4.png', 'High enjoy maintain prevent serious else suddenly.', 'Chevrolet', 'Impala', '2017',
  '1-R8201', 'nQS1109151WTTAjvd', 'automatic', 'gas', 116557,
  ST_SetSRID(ST_MakePoint(156.387010, 22.334623), 4326), 31.14, 'rented',
  '2025-04-29 20:30:41', '2025-04-29 20:30:41'
);

INSERT INTO cars (
  id, owner_id, thumbnail_picture, description, make, model, year,
  license_plate, vin_number, transmission, fuel_type, mileage, location,
  price_per_hour, status, created_at, updated_at
) VALUES (
  '10b827fc-e870-4bb1-a228-e4d6e4c71e20', '83374bb6-fd22-4d5e-89ee-b3d4faa19149', 'http://localhost:9090/static/carphoto5.png', 'National quite citizen east evidence entire.', 'Tesla', 'Model 3', '2018',
  '149XZX', 'nFW2341745HfcrxKD', 'manual', 'electric', 50723,
  ST_SetSRID(ST_MakePoint(-20.983504, -55.5818665), 4326), 10.77, 'avaliable',
  '2025-04-29 20:30:41', '2025-04-29 20:30:41'
);

INSERT INTO cars (
  id, owner_id, thumbnail_picture, description, make, model, year,
  license_plate, vin_number, transmission, fuel_type, mileage, location,
  price_per_hour, status, created_at, updated_at
) VALUES (
  '706e88b1-383c-484b-ab2b-499842223ddc', '411b85a8-586b-4089-9a34-4770494fd735', 'http://localhost:9090/static/carphoto6.png', 'Special include attack business matter central.', 'BMW', '3 Series', '2008',
  'ECY E23', 'fIb5854292cuqLpSi', 'manual', 'hybrid', 139999,
  ST_SetSRID(ST_MakePoint(-80.285497, 45.761224), 4326), 32.3, 'inactive',
  '2025-04-29 20:30:41', '2025-04-29 20:30:41'
);

INSERT INTO cars (
  id, owner_id, thumbnail_picture, description, make, model, year,
  license_plate, vin_number, transmission, fuel_type, mileage, location,
  price_per_hour, status, created_at, updated_at
) VALUES (
  'bcd975cb-c37f-4a34-a73c-d615515bc395', '582074ef-29e2-4366-9942-66afe73f070b', 'http://localhost:9090/static/carphoto7.png', 'Find rock strategy probably others together catch baby responsibility blood majority offer recognize.', 'Audi', 'A4', '2020',
  '595-RBN', 'MOj1441367HbGXzKA', 'automatic', 'hybrid', 111732,
  ST_SetSRID(ST_MakePoint(40.480487, 89.3696135), 4326), 7.32, 'rented',
  '2025-04-29 20:30:41', '2025-04-29 20:30:41'
);

INSERT INTO cars (
  id, owner_id, thumbnail_picture, description, make, model, year,
  license_plate, vin_number, transmission, fuel_type, mileage, location,
  price_per_hour, status, created_at, updated_at
) VALUES (
  '4729f1d1-2bd4-4d53-9805-34b0ff9601e4', '9306824b-c7bc-47a1-b0d0-f0b23ec36404', 'http://localhost:9090/static/carphoto8.png', 'Week station lot situation move chance father if benefit very purpose bar.', 'Mercedes', 'C-Class', '2006',
  '824 BDL', 'AxF6777767LJeMnjd', 'automatic', 'diesel', 77825,
  ST_SetSRID(ST_MakePoint(99.462505, -3.990476), 4326), 40.2, 'inactive',
  '2025-04-29 20:30:41', '2025-04-29 20:30:41'
);

INSERT INTO cars (
  id, owner_id, thumbnail_picture, description, make, model, year,
  license_plate, vin_number, transmission, fuel_type, mileage, location,
  price_per_hour, status, created_at, updated_at
) VALUES (
  '3dbc6a3d-4d7e-469f-bb02-3a761369b447', 'a8de5b66-18d5-4e42-be53-414046b680fc', 'http://localhost:9090/static/carphoto9.png', 'Heavy leg discover difference return play talk finish project take drop.', 'Volkswagen', 'Golf', '2023',
  '591580', 'jxz3471083hOVpYHr', 'automatic', 'hybrid', 106996,
  ST_SetSRID(ST_MakePoint(13.542119, -57.1903205), 4326), 19.69, 'rented',
  '2025-04-29 20:30:41', '2025-04-29 20:30:41'
);

INSERT INTO cars (
  id, owner_id, thumbnail_picture, description, make, model, year,
  license_plate, vin_number, transmission, fuel_type, mileage, location,
  price_per_hour, status, created_at, updated_at
) VALUES (
  '8028ccec-1ac6-48ef-befd-0286b88f56f5', '4c828120-5a14-4fb5-abc9-4677181251d0', 'http://localhost:9090/static/carphoto10.png', 'Meet civil before whom scene report sound son thank.', 'Hyundai', 'Elantra', '2017',
  '843 HYD', 'YEx1159167ZBuUUXz', 'manual', 'gas', 70876,
  ST_SetSRID(ST_MakePoint(28.832247, 58.619270), 4326), 17.61, 'avaliable',
  '2025-04-29 20:30:41', '2025-04-29 20:30:41'
);

INSERT INTO cars (
  id, owner_id, thumbnail_picture, description, make, model, year,
  license_plate, vin_number, transmission, fuel_type, mileage, location,
  price_per_hour, status, created_at, updated_at
) VALUES (
  'c008f474-3fb1-4276-a82d-4042dd85240e', '03ed8e1a-cedc-4641-855d-f2e94f02cd27', 'http://localhost:9090/static/carphoto11.png', 'Federal how run business send know already together paper particular doctor water.', 'Kia', 'Optima', '2012',
  'VIX 668', 'XKP8952207boeiOWK', 'manual', 'hybrid', 111014,
  ST_SetSRID(ST_MakePoint(167.356628, 4.9205215), 4326), 12.18, 'inactive',
  '2025-04-29 20:30:41', '2025-04-29 20:30:41'
);

INSERT INTO cars (
  id, owner_id, thumbnail_picture, description, make, model, year,
  license_plate, vin_number, transmission, fuel_type, mileage, location,
  price_per_hour, status, created_at, updated_at
) VALUES (
  'fb3d42b7-a0d0-43ae-9dee-8615238b43be', 'edc97d67-df03-4401-a1e2-749ff00d0297', 'http://localhost:9090/static/carphoto12.png', 'Forward safe after would treat generation him fact rise.', 'Nissan', 'Altima', '2020',
  'XXO-026', 'hkI3574041KhsMwXb', 'automatic', 'electric', 113357,
  ST_SetSRID(ST_MakePoint(-172.304026, -73.124923), 4326), 28.55, 'inactive',
  '2025-04-29 20:30:41', '2025-04-29 20:30:41'
);

INSERT INTO cars (
  id, owner_id, thumbnail_picture, description, make, model, year,
  license_plate, vin_number, transmission, fuel_type, mileage, location,
  price_per_hour, status, created_at, updated_at
) VALUES (
  'f52616a6-aabd-40aa-a86c-8ed48cfb120a', 'a3418cb9-2b0a-4f93-832c-8ee7132e5cbf', 'http://localhost:9090/static/carphoto1.png', 'Hope wrong kitchen last government time tell minute show.', 'Mazda', 'Mazda3', '2007',
  'VWQ8463', 'Hkq5750043kstLimQ', 'manual', 'diesel', 128927,
  ST_SetSRID(ST_MakePoint(124.481838, 9.382922), 4326), 19.99, 'inactive',
  '2025-04-29 20:30:41', '2025-04-29 20:30:41'
);

INSERT INTO cars (
  id, owner_id, thumbnail_picture, description, make, model, year,
  license_plate, vin_number, transmission, fuel_type, mileage, location,
  price_per_hour, status, created_at, updated_at
) VALUES (
  '14b1f196-6e45-422b-8d0f-96ed2d80bc00', '1ba44b55-56b5-49af-9396-dc9b848b4362', 'http://localhost:9090/static/carphoto2.png', 'Seem story report company person firm local responsibility skill scene when gas career.', 'Subaru', 'Impreza', '2022',
  '714 HUO', 'FSj3998588WwXlFKu', 'manual', 'gas', 87038,
  ST_SetSRID(ST_MakePoint(126.716192, -12.4537915), 4326), 20.5, 'inactive',
  '2025-04-29 20:30:41', '2025-04-29 20:30:41'
);

INSERT INTO cars (
  id, owner_id, thumbnail_picture, description, make, model, year,
  license_plate, vin_number, transmission, fuel_type, mileage, location,
  price_per_hour, status, created_at, updated_at
) VALUES (
  '982485eb-3eb9-4463-ae93-7cccd1211895', 'b9686465-ab32-43bf-9247-92008a4c1610', 'http://localhost:9090/static/carphoto3.png', 'Authority item since impact take western spring last develop result his.', 'Jeep', 'Wrangler', '2015',
  '656 KTX', 'fdj6733355kFypdAr', 'automatic', 'hybrid', 90235,
  ST_SetSRID(ST_MakePoint(-141.679168, -45.301175), 4326), 15.28, 'avaliable',
  '2025-04-29 20:30:41', '2025-04-29 20:30:41'
);

INSERT INTO cars (
  id, owner_id, thumbnail_picture, description, make, model, year,
  license_plate, vin_number, transmission, fuel_type, mileage, location,
  price_per_hour, status, created_at, updated_at
) VALUES (
  'da201407-3c8b-4ffd-af1a-4f110621b5e2', '6ce05b2f-ff7c-4574-8468-f1b68f976e86', 'http://localhost:9090/static/carphoto4.png', 'Beat current it produce similar according everybody lose.', 'Dodge', 'Charger', '2011',
  '4IM9978', 'Cii8543230CUjeRoR', 'automatic', 'electric', 94967,
  ST_SetSRID(ST_MakePoint(-84.668142, -13.6548975), 4326), 33.32, 'avaliable',
  '2025-04-29 20:30:41', '2025-04-29 20:30:41'
);

INSERT INTO cars (
  id, owner_id, thumbnail_picture, description, make, model, year,
  license_plate, vin_number, transmission, fuel_type, mileage, location,
  price_per_hour, status, created_at, updated_at
) VALUES (
  '71b830d0-b50e-4bb5-bcc3-a6a59799a47b', '9849dee0-0f04-4b5d-b6f5-d2f92f69d332', 'http://localhost:9090/static/carphoto5.png', 'Ground today receive structure owner race.', 'Chevrolet', 'Malibu', '2020',
  '811-LGZ', 'aVP9107142uNlZzCY', 'automatic', 'gas', 60346,
  ST_SetSRID(ST_MakePoint(-54.910216, 14.6486165), 4326), 7.91, 'inactive',
  '2025-04-29 20:30:41', '2025-04-29 20:30:41'
);

INSERT INTO cars (
  id, owner_id, thumbnail_picture, description, make, model, year,
  license_plate, vin_number, transmission, fuel_type, mileage, location,
  price_per_hour, status, created_at, updated_at
) VALUES (
  '69ad4016-b70d-40ea-aa11-f46f37414479', '500ea834-ab03-49cf-ac84-f5f9043d46c4', 'http://localhost:9090/static/carphoto6.png', 'Reach new international however himself hour choice begin free deep next.', 'Honda', 'Accord', '2023',
  'JDL 155', 'jdl5406843uSNLqOa', 'manual', 'hybrid', 8640,
  ST_SetSRID(ST_MakePoint(107.428103, -67.196815), 4326), 47.24, 'rented',
  '2025-04-29 20:30:41', '2025-04-29 20:30:41'
);

INSERT INTO cars (
  id, owner_id, thumbnail_picture, description, make, model, year,
  license_plate, vin_number, transmission, fuel_type, mileage, location,
  price_per_hour, status, created_at, updated_at
) VALUES (
  'a0154447-92bc-4e8c-ba65-d4fc69b7478e', '08ce3da1-775e-4dd6-93fc-67b0a546dd60', 'http://localhost:9090/static/carphoto7.png', 'Loss approach record your claim course.', 'Toyota', 'Corolla', '2024',
  '2GU60', 'lZE0309201ynkNqup', 'manual', 'electric', 60966,
  ST_SetSRID(ST_MakePoint(130.352374, -83.3121255), 4326), 37.74, 'rented',
  '2025-04-29 20:30:41', '2025-04-29 20:30:41'
);

INSERT INTO cars (
  id, owner_id, thumbnail_picture, description, make, model, year,
  license_plate, vin_number, transmission, fuel_type, mileage, location,
  price_per_hour, status, created_at, updated_at
) VALUES (
  'a6c4a0bd-7199-4543-99b9-7e7bdfe4d8a9', 'cbcf57f4-c136-4a97-9dc5-af7f71221966', 'http://localhost:9090/static/carphoto8.png', 'Medical cultural offer bit institution action.', 'Tesla', 'Model Y', '2006',
  '6RN V04', 'pFt9038940xNpWdwY', 'manual', 'hybrid', 13131,
  ST_SetSRID(ST_MakePoint(115.175055, -46.5278295), 4326), 28.59, 'inactive',
  '2025-04-29 20:30:41', '2025-04-29 20:30:41'
);

INSERT INTO cars (
  id, owner_id, thumbnail_picture, description, make, model, year,
  license_plate, vin_number, transmission, fuel_type, mileage, location,
  price_per_hour, status, created_at, updated_at
) VALUES (
  'f1bf3f1b-738d-4d47-8e4e-9ddaa1c0dc90', '7c8345f9-c839-4604-b209-a6a07ee46dde', 'http://localhost:9090/static/carphoto9.png', 'Off our whom rest official possible detail.', 'BMW', 'X5', '2020',
  'PMF3661', 'DRM7974041qrRQjqG', 'manual', 'electric', 126390,
  ST_SetSRID(ST_MakePoint(-80.577398, -9.4235225), 4326), 44.67, 'rented',
  '2025-04-29 20:30:41', '2025-04-29 20:30:41'
);

INSERT INTO cars (
  id, owner_id, thumbnail_picture, description, make, model, year,
  license_plate, vin_number, transmission, fuel_type, mileage, location,
  price_per_hour, status, created_at, updated_at
) VALUES (
  'eb82b42a-b9d0-4f72-92ee-eda1348109fb', '32b30c74-06c1-4644-9e69-b5b2094cd298', 'http://localhost:9090/static/carphoto10.png', 'Matter work however low hospital edge result billion itself.', 'Audi', 'Q5', '2017',
  'JUU 120', 'PwE0052185kIDeMIp', 'manual', 'electric', 29678,
  ST_SetSRID(ST_MakePoint(16.275744, -2.439424), 4326), 43.95, 'rented',
  '2025-04-29 20:30:41', '2025-04-29 20:30:41'
);

INSERT INTO cars (
  id, owner_id, thumbnail_picture, description, make, model, year,
  license_plate, vin_number, transmission, fuel_type, mileage, location,
  price_per_hour, status, created_at, updated_at
) VALUES (
  '0dcc885f-be73-4886-ab17-6e65ea426474', '4f1e3215-5576-4127-8a15-22a8f6492d96', 'http://localhost:9090/static/carphoto11.png', 'Sign will oil experience president bar.', 'Ford', 'Explorer', '2012',
  '355 TSS', 'vpn4949872dzOjVEz', 'automatic', 'electric', 5065,
  ST_SetSRID(ST_MakePoint(156.216438, -38.162300), 4326), 41.46, 'avaliable',
  '2025-04-29 20:30:41', '2025-04-29 20:30:41'
);

INSERT INTO cars (
  id, owner_id, thumbnail_picture, description, make, model, year,
  license_plate, vin_number, transmission, fuel_type, mileage, location,
  price_per_hour, status, created_at, updated_at
) VALUES (
  'e883d57b-fa11-4892-8c99-9e0c91aa01fc', 'f8b859c5-972a-4569-999a-b49225f09b48', 'http://localhost:9090/static/carphoto12.png', 'Maybe but teach against themselves single rich wife example use writer magazine.', 'Nissan', 'Rogue', '2022',
  'IWY-528', 'dxG1494131IYIyXSR', 'manual', 'diesel', 39064,
  ST_SetSRID(ST_MakePoint(81.485629, -7.5230205), 4326), 34.28, 'inactive',
  '2025-04-29 20:30:41', '2025-04-29 20:30:41'
);

INSERT INTO cars (
  id, owner_id, thumbnail_picture, description, make, model, year,
  license_plate, vin_number, transmission, fuel_type, mileage, location,
  price_per_hour, status, created_at, updated_at
) VALUES (
  '43ce82e0-7632-4e45-85f3-365f0aefb8d8', 'f8e5eda9-a3c4-4026-a549-7fed2816bcb3', 'http://localhost:9090/static/carphoto1.png', 'Dog thing they PM without shake more national movie kitchen different nothing.', 'Kia', 'Sorento', '2024',
  '07J O34', 'FOV3814022yUikqlF', 'manual', 'gas', 145717,
  ST_SetSRID(ST_MakePoint(-98.879792, -46.422463), 4326), 12.11, 'inactive',
  '2025-04-29 20:30:41', '2025-04-29 20:30:41'
);

INSERT INTO cars (
  id, owner_id, thumbnail_picture, description, make, model, year,
  license_plate, vin_number, transmission, fuel_type, mileage, location,
  price_per_hour, status, created_at, updated_at
) VALUES (
  'ba44ec3a-38b5-4eb5-8fc2-1e4c7260f1c9', '544139f5-395d-4465-96fa-15a3c8db7bcc', 'http://localhost:9090/static/carphoto2.png', 'Material Congress audience common worry try yes contain family all up save.', 'Toyota', 'Camry', '2008',
  'WVO-1015', 'Xpp0825844jgNhLie', 'automatic', 'hybrid', 20162,
  ST_SetSRID(ST_MakePoint(70.365796, 44.2583175), 4326), 27.68, 'avaliable',
  '2025-04-29 20:30:41', '2025-04-29 20:30:41'
);

INSERT INTO cars (
  id, owner_id, thumbnail_picture, description, make, model, year,
  license_plate, vin_number, transmission, fuel_type, mileage, location,
  price_per_hour, status, created_at, updated_at
) VALUES (
  '743da0a9-bb90-4267-95ee-d9d3f4ce2343', 'b8a35aed-5a8b-4c85-ae5d-55867c2b02cf', 'http://localhost:9090/static/carphoto3.png', 'Child she public only likely loss cost step.', 'Honda', 'Civic', '2022',
  '938HNM', 'woM9329665gjOFDqE', 'automatic', 'gas', 9719,
  ST_SetSRID(ST_MakePoint(-147.267854, -61.4813735), 4326), 32.05, 'inactive',
  '2025-04-29 20:30:41', '2025-04-29 20:30:41'
);

INSERT INTO cars (
  id, owner_id, thumbnail_picture, description, make, model, year,
  license_plate, vin_number, transmission, fuel_type, mileage, location,
  price_per_hour, status, created_at, updated_at
) VALUES (
  '5e9987fd-dfa1-481a-bf5d-f870842cf7bf', '1ae96d1d-0ddd-41a8-bb8c-c8b349c39b50', 'http://localhost:9090/static/carphoto4.png', 'Soon many base current where your early.', 'Ford', 'Mustang', '2021',
  'LCE 760', 'PjI5908862puJaYgL', 'automatic', 'hybrid', 40490,
  ST_SetSRID(ST_MakePoint(-142.830922, -46.011576), 4326), 7.83, 'inactive',
  '2025-04-29 20:30:41', '2025-04-29 20:30:41'
);

INSERT INTO cars (
  id, owner_id, thumbnail_picture, description, make, model, year,
  license_plate, vin_number, transmission, fuel_type, mileage, location,
  price_per_hour, status, created_at, updated_at
) VALUES (
  '23d678fe-87b9-4d43-bd3d-177728d82efc', 'd21e5eaf-674b-41bc-a0ec-2e75cf7ff33f', 'http://localhost:9090/static/carphoto5.png', 'Before measure even try wrong head firm campaign return per enough.', 'Chevrolet', 'Impala', '2008',
  'LNT7245', 'tjK4712987hmxKxKQ', 'manual', 'gas', 149424,
  ST_SetSRID(ST_MakePoint(-32.471805, 9.0383805), 4326), 10.56, 'rented',
  '2025-04-29 20:30:41', '2025-04-29 20:30:41'
);

INSERT INTO cars (
  id, owner_id, thumbnail_picture, description, make, model, year,
  license_plate, vin_number, transmission, fuel_type, mileage, location,
  price_per_hour, status, created_at, updated_at
) VALUES (
  '25c2217b-e3a9-4cca-838f-c2cf21e24246', 'b12e02da-dc66-409a-84ba-b2f4def6e818', 'http://localhost:9090/static/carphoto6.png', 'Trouble bag than deep prevent trouble why respond exactly role teacher sing.', 'Tesla', 'Model 3', '2015',
  '6-2317S', 'tnj8961067mCWCIAh', 'automatic', 'diesel', 38081,
  ST_SetSRID(ST_MakePoint(-145.314304, -26.6999895), 4326), 43.73, 'avaliable',
  '2025-04-29 20:30:41', '2025-04-29 20:30:41'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE cars CASCADE;
-- +goose StatementEnd
