CREATE TABLE mst_service
(
    id character varying(100) NOT NULL,
    pelayanan character varying(100) NOT NULL,
    harga integer NOT NULL,
    CONSTRAINT mst_table_pkey PRIMARY KEY (no)
);

CREATE TABLE mst_customer
(
    id character varying(100) NOT NULL,
    nama_customer character varying(50) NOT NULL,
    nomor_hp character varying(15) NOT NULL,
);

CREATE TABLE trx_bill
(
    id integer NOT NULL,
    tanggal_masuk date NOT NULL,
    tanggal_selesai date NOT NULL,
    diterima_oleh character varying(100) NOT NULL,
    customer_id character varying(100) NOT NULL,
    CONSTRAINT mst_table_pkey PRIMARY KEY (no)
);

ALTER table trx_bill ADD CONSTRAINT fk_trx_bill foreign key (customer_id) references mst_customer(id)

CREATE TABLE trx_bill_detail
(
    no integer NOT NULL,
    bill_id character varying(100) NOT NULL,
    pelayanan_id character varying(100) NOT NULL,
    jumlah integer NOT NULL,
    satuan character varying(15) NOT NULL,
    harga integer NOT NULL,
    total integer NOT NULL,
    CONSTRAINT trx_bill_pkey PRIMARY KEY (no)
);

ALTER table trx_bill_detail ADD CONSTRAINT fk_trx_bill_detail foreign key (bill_id) references trx_bill(id)

ALTER table trx_bill_detail ADD CONSTRAINT fk_trx_bill_detail_pelayanan foreign key (pelayanan_id) references mst_service(id)

