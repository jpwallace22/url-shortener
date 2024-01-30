-- CreateTable
CREATE TABLE "User" (
    "id" SERIAL NOT NULL,
    "url" TEXT NOT NULL,
    "hash" TEXT NOT NULL,
    "locked" BOOLEAN DEFAULT false,
    "password" TEXT,

    CONSTRAINT "User_pkey" PRIMARY KEY ("id")
);
