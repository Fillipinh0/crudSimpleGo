CREATE DATABASE IF NOT EXISTS db_trabalho;

USE db_trabalho;

CREATE TABLE IF NOT EXISTS cliente (
    id_cliente INT PRIMARY KEY auto_increment,
    nome_cliente VARCHAR(70),
    telefone_cliente VARCHAR(15),
    renda_cliente VARCHAR(15),
    data_nascimento_cliente DATE,
    CEP_cliente VARCHAR(10)

)