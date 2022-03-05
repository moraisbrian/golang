USE master;

GO

CREATE DATABASE GoTeste;

GO

USE [GoTeste]

GO

CREATE TABLE [dbo].[Reminders](
	[ID] [int] IDENTITY(1,1) PRIMARY KEY NOT NULL,
	[Title] [varchar](70) NULL,
	[Description] [varchar](70) NULL,
	[Alias] [varchar](70) NULL
);

GO