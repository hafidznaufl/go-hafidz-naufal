-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Sep 10, 2023 at 02:02 PM
-- Server version: 10.4.27-MariaDB
-- PHP Version: 8.2.0

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `tugas`
--

DELIMITER $$
--
-- Functions
--
CREATE DEFINER=`root`@`localhost` FUNCTION `delete_transaction_details` (`transaction_id` INT) RETURNS TINYINT(1)  BEGIN
    DELETE FROM transaction_details WHERE transaction_id = transaction_id;
    RETURN TRUE;
END$$

CREATE DEFINER=`root`@`localhost` FUNCTION `update_total_qty` (`transaction_id` INT) RETURNS TINYINT(1)  BEGIN
    DECLARE deleted_qty INT;
    
    -- Menghitung jumlah qty yang dihapus
    SELECT SUM(qty) INTO deleted_qty
    FROM transaction_details
    WHERE transaction_id = transaction_id;
    
    -- Mengurangkan jumlah qty yang dihapus dari total_qty
    UPDATE transactions
    SET total_qty = total_qty - deleted_qty
    WHERE id = transaction_id;
    
    RETURN TRUE;
END$$

DELIMITER ;

-- --------------------------------------------------------

--
-- Table structure for table `operators`
--

CREATE TABLE `operators` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `operators`
--

INSERT INTO `operators` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'Ahmad', '2023-09-09 08:28:27', '2023-09-09 08:28:27'),
(2, 'Boro', '2023-09-09 08:29:38', '2023-09-09 08:29:38'),
(3, 'Oda', '2023-09-09 08:29:53', '2023-09-09 08:29:53'),
(4, 'Iwan', '2023-09-09 08:30:33', '2023-09-09 08:30:33'),
(5, 'Kina', '2023-09-09 08:30:44', '2023-09-09 08:30:44');

-- --------------------------------------------------------

--
-- Table structure for table `payment_methods`
--

CREATE TABLE `payment_methods` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `status` smallint(6) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `payment_methods`
--

INSERT INTO `payment_methods` (`id`, `name`, `status`, `created_at`, `updated_at`) VALUES
(1, 'Bank', 1, '2023-09-09 09:10:58', '2023-09-09 09:10:58'),
(2, 'GOpay', 1, '2023-09-09 09:11:07', '2023-09-09 09:11:07'),
(3, 'Ovo', 1, '2023-09-09 09:11:19', '2023-09-09 09:11:19');

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `id` int(11) NOT NULL,
  `product_types_id` int(11) NOT NULL,
  `operators_id` int(11) NOT NULL,
  `product_descriptions_id` int(11) NOT NULL,
  `code` varchar(50) NOT NULL,
  `name` varchar(100) DEFAULT NULL,
  `status` smallint(6) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `products`
--

INSERT INTO `products` (`id`, `product_types_id`, `operators_id`, `product_descriptions_id`, `code`, `name`, `status`, `created_at`, `updated_at`) VALUES
(2, 2, 3, 1, '002', 'Majalah Vogue', 18, '2023-09-09 08:45:26', '2023-09-09 09:48:42'),
(3, 2, 1, 2, '003', 'Pensil', 29, '2023-09-09 08:46:22', '2023-09-09 09:08:55'),
(4, 2, 1, 2, '004', 'Pulpen', 21, '2023-09-09 08:46:40', '2023-09-09 09:08:49'),
(5, 2, 1, 2, '005', 'Crayon', 34, '2023-09-09 08:47:06', '2023-09-09 09:09:09'),
(6, 3, 4, 3, '006', 'Penggaris', 23, '2023-09-09 08:48:09', '2023-09-09 09:09:14'),
(7, 3, 4, 3, '007', 'Jangka', 29, '2023-09-09 08:48:55', '2023-09-09 09:09:20'),
(8, 3, 4, 3, '008', 'Neraca', 12, '2023-09-09 08:49:16', '2023-09-09 09:09:26');

-- --------------------------------------------------------

--
-- Table structure for table `product_descriptions`
--

CREATE TABLE `product_descriptions` (
  `id` int(11) NOT NULL,
  `description` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `product_descriptions`
--

INSERT INTO `product_descriptions` (`id`, `description`, `created_at`, `updated_at`) VALUES
(1, 'Buku', '2023-09-09 09:07:54', '2023-09-09 09:07:54'),
(2, 'Alat Tulis', '2023-09-09 09:08:03', '2023-09-09 09:08:03'),
(3, 'Alat Ukur', '2023-09-09 09:08:10', '2023-09-09 09:08:10');

-- --------------------------------------------------------

--
-- Table structure for table `product_types`
--

CREATE TABLE `product_types` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `product_types`
--

INSERT INTO `product_types` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'Buku', '2023-09-09 08:40:38', '2023-09-09 08:40:38'),
(2, 'Alat Tulis', '2023-09-09 08:40:47', '2023-09-09 08:40:47'),
(3, 'Alat Ukur', '2023-09-09 08:41:12', '2023-09-09 08:41:12');

-- --------------------------------------------------------

--
-- Table structure for table `transactions`
--

CREATE TABLE `transactions` (
  `id` int(11) NOT NULL,
  `users_id` int(11) DEFAULT NULL,
  `payment_methods_id` int(11) DEFAULT NULL,
  `status` varchar(10) DEFAULT NULL,
  `total_qty` int(11) DEFAULT NULL,
  `total_price` decimal(24,2) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `transactions`
--

INSERT INTO `transactions` (`id`, `users_id`, `payment_methods_id`, `status`, `total_qty`, `total_price`, `created_at`, `updated_at`) VALUES
(1, 1, 1, '1', 3, '20000.00', '2023-09-09 09:14:11', '2023-09-09 09:42:32'),
(2, 2, 2, '1', 31, '20000.00', '2023-09-09 09:14:31', '2023-09-09 09:14:31'),
(3, 3, 3, '1', 20, '30000.00', '2023-09-09 09:14:48', '2023-09-09 09:14:48');

-- --------------------------------------------------------

--
-- Table structure for table `transaction_details`
--

CREATE TABLE `transaction_details` (
  `id` int(11) NOT NULL,
  `transaction_id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  `dob` date NOT NULL,
  `gender` char(1) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `transaction_details`
--

INSERT INTO `transaction_details` (`id`, `transaction_id`, `product_id`, `dob`, `gender`, `created_at`, `updated_at`) VALUES
(1, 2, 4, '2023-09-18', 'L', '2023-09-09 09:15:37', '2023-09-09 09:15:37'),
(2, 1, 2, '2023-09-18', 'P', '2023-09-09 09:15:46', '2023-09-09 09:15:46'),
(3, 2, 4, '2023-09-18', 'L', '2023-09-09 09:15:54', '2023-09-09 09:15:54');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `name` varchar(50) DEFAULT NULL,
  `status` smallint(6) NOT NULL,
  `dob` date NOT NULL,
  `gender` char(1) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `status`, `dob`, `gender`, `created_at`, `updated_at`) VALUES
(1, 'FDQ', 1, '2023-09-07', 'L', '2023-09-09 09:12:40', '2023-09-09 09:28:22'),
(2, 'WEQ', 1, '2023-09-07', 'L', '2023-09-09 09:12:49', '2023-09-09 09:28:26'),
(3, 'REW', 1, '2023-09-07', 'L', '2023-09-09 09:12:55', '2023-09-09 09:28:31'),
(4, 'ASDAW', 1, '2023-09-07', 'L', '2023-09-09 09:13:03', '2023-09-09 09:28:37'),
(5, 'ASDW', 1, '2023-09-07', 'P', '2023-09-09 09:13:11', '2023-09-09 09:28:41'),
(6, 'FFGSD', 1, '2023-09-07', 'P', '2023-09-09 09:13:17', '2023-09-09 09:28:45');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `operators`
--
ALTER TABLE `operators`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `payment_methods`
--
ALTER TABLE `payment_methods`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`),
  ADD KEY `product_types_id` (`product_types_id`),
  ADD KEY `operators_id` (`operators_id`),
  ADD KEY `product_descriptions` (`product_descriptions_id`);

--
-- Indexes for table `product_descriptions`
--
ALTER TABLE `product_descriptions`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `product_types`
--
ALTER TABLE `product_types`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `users_id` (`users_id`),
  ADD KEY `payment_methods_id` (`payment_methods_id`);

--
-- Indexes for table `transaction_details`
--
ALTER TABLE `transaction_details`
  ADD PRIMARY KEY (`id`),
  ADD KEY `transaction_id` (`transaction_id`),
  ADD KEY `product_id` (`product_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `operators`
--
ALTER TABLE `operators`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `payment_methods`
--
ALTER TABLE `payment_methods`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `products`
--
ALTER TABLE `products`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `product_descriptions`
--
ALTER TABLE `product_descriptions`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `product_types`
--
ALTER TABLE `product_types`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `transactions`
--
ALTER TABLE `transactions`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `transaction_details`
--
ALTER TABLE `transaction_details`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `products`
--
ALTER TABLE `products`
  ADD CONSTRAINT `fk_product_description` FOREIGN KEY (`product_descriptions_id`) REFERENCES `product_descriptions` (`id`),
  ADD CONSTRAINT `product_descriptions` FOREIGN KEY (`product_descriptions_id`) REFERENCES `product_descriptions` (`id`),
  ADD CONSTRAINT `products_ibfk_1` FOREIGN KEY (`product_types_id`) REFERENCES `product_types` (`id`),
  ADD CONSTRAINT `products_ibfk_2` FOREIGN KEY (`operators_id`) REFERENCES `operators` (`id`);

--
-- Constraints for table `transactions`
--
ALTER TABLE `transactions`
  ADD CONSTRAINT `transactions_ibfk_1` FOREIGN KEY (`users_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `transactions_ibfk_2` FOREIGN KEY (`payment_methods_id`) REFERENCES `payment_methods` (`id`);

--
-- Constraints for table `transaction_details`
--
ALTER TABLE `transaction_details`
  ADD CONSTRAINT `transaction_details_ibfk_1` FOREIGN KEY (`transaction_id`) REFERENCES `transactions` (`id`),
  ADD CONSTRAINT `transaction_details_ibfk_2` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
