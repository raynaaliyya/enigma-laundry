INSERT INTO mst_service(
	id, pelayanan, harga)
	VALUES ('S001', 'Cuci + Setrika', 7000);

INSERT INTO mst_customer(
	id, nama_customer, nomor_hp)
	VALUES ('C001', 'Jessica', '0812654987');

INSERT INTO trx_bill(
	id, tanggal_masuk, tanggal_selesai, diterima_oleh, customer_id)
	VALUES ('B001', '18/08/2022', '20/08/2022', 'Mirna', 'C001');

INSERT INTO trx_bill_detail(
	id, bill_id, pelayanan_id, jumlah, satuan, harga, total)
	VALUES ('BD001', 'B001', 'S001', 5, 'KG', 7000, 35000);

