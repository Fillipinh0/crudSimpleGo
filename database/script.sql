

CREATE DATABASE IF NOT EXISTS db_trabalho;

USE db_trabalho;

CREATE TABLE IF NOT EXISTS ativo (
    id_ativo INT PRIMARY KEY auto_increment,
    nome_ativo VARCHAR(50),
    localizacao_ativo VARCHAR(50),
    descricao_ativo VARCHAR(150),
    status_ativo VARCHAR(10)
)