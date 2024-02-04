/*
  Warnings:

  - The primary key for the `Url` table will be changed. If it partially fails, the table could be left without primary key constraint.
  - You are about to drop the column `id` on the `Url` table. All the data in the column will be lost.
  - You are about to drop the column `shortUrl` on the `Url` table. All the data in the column will be lost.
  - Added the required column `short_url` to the `Url` table without a default value. This is not possible if the table is not empty.
  - Added the required column `url_id` to the `Url` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "Url" DROP CONSTRAINT "Url_pkey",
DROP COLUMN "id",
DROP COLUMN "shortUrl",
ADD COLUMN     "short_url" TEXT NOT NULL,
ADD COLUMN     "url_id" TEXT NOT NULL,
ADD CONSTRAINT "Url_pkey" PRIMARY KEY ("url_id");
