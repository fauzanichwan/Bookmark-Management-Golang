-- phpMyAdmin SQL Dump
-- version 4.7.4
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jun 05, 2018 at 08:55 AM
-- Server version: 10.1.30-MariaDB
-- PHP Version: 7.2.1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `bookmark_db`
--

-- --------------------------------------------------------

--
-- Table structure for table `admin`
--

CREATE TABLE `admin` (
  `idAdmin` int(11) NOT NULL,
  `username` varchar(50) NOT NULL,
  `password` varchar(255) NOT NULL,
  `foto` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `berkas`
--

CREATE TABLE `berkas` (
  `idBerkas` int(11) NOT NULL,
  `idUser` int(11) NOT NULL,
  `namaBerkas` varchar(50) NOT NULL,
  `berkas` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `bookmark`
--

CREATE TABLE `bookmark` (
  `idBookmark` int(11) NOT NULL,
  `nmKategori` varchar(50) NOT NULL,
  `idUser` int(11) NOT NULL,
  `judul` varchar(20) NOT NULL,
  `link` varchar(255) NOT NULL,
  `tglBuat` varchar(30) NOT NULL,
  `status` int(1) NOT NULL DEFAULT '0',
  `publish` int(1) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `bookmark`
--

INSERT INTO `bookmark` (`idBookmark`, `nmKategori`, `idUser`, `judul`, `link`, `tglBuat`, `status`, `publish`) VALUES
(4, 'forum', 16, 'forum famous', 'https://kaskus.co.id', '', 0, 1),
(5, 'rpl', 17, 'website sekolah ku', 'https://dwiguna.co.id', '', 1, 0),
(6, 'sepak bola', 18, 'website fifa', 'https://fifa.com', '', 0, 0),
(7, 'web', 16, 'css developer', 'https://w3schools.com', '', 0, 0),
(8, 'sains', 16, 'jurnal', 'https://nekonime.tv', '', 1, 1),
(9, 'action figure', 16, 'good looking', 'https://facebook.com', '', 0, 0),
(10, 'golang', 17, 'website resminya', 'https://golang.org', '', 1, 0),
(11, 'rpl', 17, 'forumnya bro', 'https://www.youtube.com/watch?v=Yg3LxxDpn84', '', 0, 0),
(12, 'css', 16, 'good dabes ubah', 'https://youtube.com', '', 1, 0),
(13, 'game', 19, 'pc game low', 'https://bagas31.com', '', 0, 0),
(15, 'forum', 16, 'ketiga ubah', 'https://student.gunadarma.ac.id', '', 0, 0);

-- --------------------------------------------------------

--
-- Table structure for table `kategori`
--

CREATE TABLE `kategori` (
  `idKategori` int(11) NOT NULL,
  `idUser` int(11) NOT NULL,
  `nmKategori` varchar(25) NOT NULL,
  `tglBuat` varchar(30) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `kategori`
--

INSERT INTO `kategori` (`idKategori`, `idUser`, `nmKategori`, `tglBuat`) VALUES
(10, 16, 'forum', ''),
(11, 16, 'web', ''),
(12, 16, 'typing', ''),
(13, 16, 'css', ''),
(14, 17, 'rpl', ''),
(15, 17, 'oop', ''),
(16, 16, 'sains', ''),
(17, 17, 'php', ''),
(18, 17, 'golang', ''),
(19, 16, 'dunia', ''),
(20, 18, 'sepak bola', ''),
(21, 18, 'training', ''),
(22, 16, 'action figure', ''),
(23, 19, 'game', ''),
(24, 19, 'bermain', ''),
(25, 19, 'olahraga', ''),
(26, 16, 'testing', '');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `idUser` int(11) NOT NULL,
  `nmDepan` varchar(50) NOT NULL,
  `nmBelakang` varchar(50) NOT NULL,
  `email` varchar(30) NOT NULL,
  `jk` varchar(6) NOT NULL,
  `username` varchar(50) NOT NULL,
  `password` varchar(255) NOT NULL,
  `tglGabung` varchar(30) NOT NULL,
  `foto` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`idUser`, `nmDepan`, `nmBelakang`, `email`, `jk`, `username`, `password`, `tglGabung`, `foto`) VALUES
(16, 'caca', 'handika fadlan', 'serginho@yahoo.com', 'pria', 'caca', 'a16358be6e2306b153b1f071477e68837266075e', '', ''),
(17, 'siska', 'suitomo', 'siska@gmail.com', 'wanita', 'siska', 'b86cd626dbdf007914168868092a30640aedffff', '', ''),
(18, 'shahril', 'ishak', 'shahril@gmail.com', 'pria', 'shahril', '0956be0deb10a16f1cd7c73279c111a3cbb182d6', '', ''),
(19, 'fauzan ichwan', 'noor pratama', 'fauzanichwan99@yahoo.co.id', 'pria', 'fauzan', '39adb8d1fba094029b87edf3f794b2e8dca196de', '', '');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `admin`
--
ALTER TABLE `admin`
  ADD PRIMARY KEY (`idAdmin`);

--
-- Indexes for table `berkas`
--
ALTER TABLE `berkas`
  ADD PRIMARY KEY (`idBerkas`);

--
-- Indexes for table `bookmark`
--
ALTER TABLE `bookmark`
  ADD PRIMARY KEY (`idBookmark`);

--
-- Indexes for table `kategori`
--
ALTER TABLE `kategori`
  ADD PRIMARY KEY (`idKategori`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`idUser`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `admin`
--
ALTER TABLE `admin`
  MODIFY `idAdmin` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `berkas`
--
ALTER TABLE `berkas`
  MODIFY `idBerkas` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `bookmark`
--
ALTER TABLE `bookmark`
  MODIFY `idBookmark` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;

--
-- AUTO_INCREMENT for table `kategori`
--
ALTER TABLE `kategori`
  MODIFY `idKategori` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=27;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `idUser` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=20;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
