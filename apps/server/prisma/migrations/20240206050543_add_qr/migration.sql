/*
  Warnings:

  - You are about to drop the column `updated_att` on the `Url` table. All the data in the column will be lost.
  - Added the required column `qr_code` to the `Url` table without a default value. This is not possible if the table is not empty.
  - Added the required column `updated_at` to the `Url` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "Url" DROP COLUMN "updated_att",
ADD COLUMN     "qr_code" TEXT NOT NULL,
ADD COLUMN     "updated_at" TIMESTAMP(3) NOT NULL;
