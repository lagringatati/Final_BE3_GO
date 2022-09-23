SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE="TRADITIONAL";

DROP SCHEMA IF EXISTS my_db;
CREATE SCHEMA my_db;
USE my_db;

--
-- Table structure for table `odontologos`
--
DROP TABLE IF EXISTS odontologos;
CREATE TABLE odontologos (
  idOdontologo INT UNSIGNED NOT NULL AUTO_INCREMENT,
  nombreOdontologo VARCHAR(50) NOT NULL,
  apellidoOdontologo VARCHAR(50) NOT NULL,
  matriculaOdontologo VARCHAR(50) NOT NULL,
  PRIMARY KEY (idOdontologo)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Table structure for table `pacientes`
--
DROP TABLE IF EXISTS pacientes;
CREATE TABLE pacientes (
  idPaciente INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT "id de la tabla pacientes",
  nombrePaciente varchar(50) NOT NULL,
  apellidoPaciente varchar(50) NOT NULL,
  domicilioPaciente varchar(50) NOT NULL,
  dniPaciente varchar(50) NOT NULL,
  fechaDeAltaPaciente datetime NOT NULL,
  PRIMARY KEY (idPaciente)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Table structure for table `turnos`
--
DROP TABLE IF EXISTS turnos;
CREATE TABLE turnos (
  idTurno INT UNSIGNED NOT NULL AUTO_INCREMENT,
  descripcionTurno VARCHAR(150) NOT NULL,
  fechaTurno DATETIME NOT NULL,
  idOdontologo INT UNSIGNED NOT NULL,
  idPaciente INT UNSIGNED NOT NULL,
  PRIMARY KEY (idTurno),
  KEY fk_odontologos_turnos(idOdontologo),
  CONSTRAINT fk_odontologos_turnos FOREIGN KEY (idOdontologo) REFERENCES odontologos(idOdontologo)  
  ON DELETE CASCADE
  ON UPDATE CASCADE,
  KEY fk_pacientes_turnos(idPaciente),
  CONSTRAINT fk_pacientes_turnos FOREIGN KEY (idPaciente) REFERENCES pacientes(idPaciente)  
  ON DELETE CASCADE
  ON UPDATE CASCADE
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `pacientes`
--
SET AUTOCOMMIT=0;
INSERT INTO pacientes(nombrePaciente, apellidoPaciente, domicilioPaciente, dniPaciente, fechaAltaPaciente) VALUES ("Cintia","Vettorazzi","Ameghino 491", "33151264", "1992-01-19");
INSERT INTO pacientes(nombrePaciente, apellidoPaciente, domicilioPaciente, dniPaciente, fechaAltaPaciente) VALUES ("Alicia","Alvarez","Juan Vicario 160", "12047906", "1987-10-08");
INSERT INTO pacientes(nombrePaciente, apellidoPaciente, domicilioPaciente, dniPaciente, fechaAltaPaciente) VALUES ("Yesica","Alvarado","Mujica Lainez 1773", "33183656", "2020-01-01");
INSERT INTO pacientes(nombrePaciente, apellidoPaciente, domicilioPaciente, dniPaciente, fechaAltaPaciente) VALUES ("Natalia","Arenas","Belgrano 502", "26147852", "2021-09-12");
COMMIT;

--
-- Dumping data for table `odontologos`
--

SET AUTOCOMMIT=0;
INSERT INTO odontologos(nombreOdontologo,apellidoOdontologo,matriculaOdontologo) VALUES ("Sergio","Macor","MP3142");
INSERT INTO odontologos(nombreOdontologo,apellidoOdontologo,matriculaOdontologo) VALUES ("Manuel","Rodriguez","MP0050");
INSERT INTO odontologos(nombreOdontologo,apellidoOdontologo,matriculaOdontologo) VALUES ("Juan","Perez","MP1234");
COMMIT;

--
-- Dumping data for table `turnos`
--

SET AUTOCOMMIT=0;
INSERT INTO turnos(descripcionTurno, fechaTurno, idOdontologo, idPaciente) VALUES ("Revision general","2022-09-23 10:00:00",1,1);
INSERT INTO turnos(descripcionTurno, fechaTurno, idOdontologo, idPaciente) VALUES ("Control","2022-09-23 10:45:00",1,2);
INSERT INTO turnos(descripcionTurno, fechaTurno, idOdontologo, idPaciente) VALUES ("Extracción","2022-09-27 18:00:00",2,3);
INSERT INTO turnos(descripcionTurno, fechaTurno, idOdontologo, idPaciente) VALUES ("Blanqueamiento","2022-09-30 16:30:00",3,4);
COMMIT;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
