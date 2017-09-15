---
title: What is Symbolic Links and how to use it?
date: '2017-08-18 11:00:00'
comments: true
categories: programming
keywords: Symbolic Links, linux, unix, terminal, hard link, link
excerpt: Alternatively referred to as a soft link or symlink, a symbolic link is a file that links to another file or directory using its path.
---

Alternatively referred to as a soft link or symlink, a symbolic link is a file that links to another file or directory using its path. Unlike a hard link, a symbolic link can link to any file or directory on any computer. In Linux and Unix symbolic links are created with the ln command, and in the Windows command line, symbolic links are created using the mklink command. Below is an example of a symbolic link in the Windows command line.

Alternatively known as SYLK, a symbolic link is an ASCII formatted file with the extension .slk, used by some Microsoft applications. It can exchange data between some Microsoft applications such as Excel.

## What Is A Hard Link

Each file in your file system is identified by a number called an inode.


Most of the time you won't really care about this but the importance of this comes to light when you want to create a hard link.

A hard link lets you assign a different name to a file in a different location but essentially it is exactly the same file. The key that links the files together is the inode number.

The great thing about hard links is that they don't take up any physical hard drive space.

A hard link makes it easier to categorize files. For instance, imagine you have a folder full of photos. You could create one folder called holiday pictures, another folder called kids photos and a third called pet photos.

It is possible that you will have some photos that fit into all three categories because they were taken on holiday with your children and dogs present.

You could put the main file in the holiday pictures photos and then create a hard link to that photo in the kid's photos category and another hard link in the pet photos category.

No extra space is taken up.

All you have to do is enter the following command to create a hard link:

```bash
ln /path/to/file /path/to/hardlink
```

Imagine you had a photo called BrightonBeach in the holiday photos folder and you wanted to create a link in the kid's photos folder you would use the following command

```bash
ln /holidayphotos/BrightonBeach.jpg /kidsphotos/BrightonBeach.jpg
```

You can tell how many files link to the same inode by using the ls command as follows: 

```bash
s -lt
```

The output will be something like -rw-r--r-- 1 username groupname date filename.

The first part shows the user's permissions. The important bit is the number after the permissions and before the username. 

If the number is 1 it is the only file pointing to a particular inode (i.e. it is not linked). If the number is greater than one then it is hard linked by 2 or more files.

## What Is A Symbolic Link

A symbolic link is like a shortcut from one file to another. The contents of a symbolic link are the address of the actual file or folder that is being linked to.

The benefit of using symbolic links is that you can link to files and folders on other partitions and on other devices.

Another difference between a hard link and a symbolic link is that a hard link must be created against a file that already exists whereas a soft link can be created in advance of the file it is pointing to existing.

To create a symbolic link use the following syntax:

```bash
ln -s /path/to/file /path/to/link
```

If you are worried about overwriting a link that already exists you can use the -b switch as follows:

```bash
ln -s -b /path/to/file /path/to/link
```

This will create a backup of the link if it already exists by creating the same filename but with a tilde at the end (~).

If a file already exists with the same name as the symbolic link you will receive an error.

You can force the link to overwrite the file by using the following command:

```bash
ln -s -f /path/to/file /path/to/link
```

You probably don't want to use the -f switch without the -b switch as you will lose the original file.

Another alternative is to receive a message asking whether you want to overwrite a file if it already exists. You can do this with the following command:

```bash
ln -s -i /path/to/file /path/to/link
```

How do you tell if a file is a symbolic link?

Run the following ls command:

```bash
s -lt
```

If a file is a symbolic link you will see something like this:

```bash
myshortcut -> myfile
```

You can use a symbolic link to navigate to another folder.

For example, imagine you have a link to /home/music/rock/alicecooper/heystoopid called heystoopid

You can run the following cd command to navigate to that folder using the following command:

```bash
cd heystoopid
```

## Summary

So that is it. You use symbolic links like shortcuts. They can be used to make really long paths shorter and a way to get easy access to files on other partitions and drives.

This guide shows everything you need to know about symbolic links but you can check out the manual page for the ln command for the other switches.