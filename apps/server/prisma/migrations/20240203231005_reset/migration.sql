-- CreateTable
CREATE TABLE "Url" (
    "id" SERIAL NOT NULL,
    "date" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_att" TIMESTAMP(3) NOT NULL,
    "shortUrl" TEXT NOT NULL,
    "url" TEXT NOT NULL,
    "clicks" INTEGER NOT NULL DEFAULT 0,
    "hash" TEXT,

    CONSTRAINT "Url_pkey" PRIMARY KEY ("id")
);
