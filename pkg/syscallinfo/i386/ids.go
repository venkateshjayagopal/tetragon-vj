// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Tetragon

package i386

// This file was generated by dump-syscalls-info

const (
	SYS_RESTART_SYSCALL              = 0
	SYS_EXIT                         = 1
	SYS_FORK                         = 2
	SYS_READ                         = 3
	SYS_WRITE                        = 4
	SYS_OPEN                         = 5
	SYS_CLOSE                        = 6
	SYS_WAITPID                      = 7
	SYS_CREAT                        = 8
	SYS_LINK                         = 9
	SYS_UNLINK                       = 10
	SYS_EXECVE                       = 11
	SYS_CHDIR                        = 12
	SYS_TIME                         = 13
	SYS_MKNOD                        = 14
	SYS_CHMOD                        = 15
	SYS_LCHOWN                       = 16
	SYS_BREAK                        = 17
	SYS_OLDSTAT                      = 18
	SYS_LSEEK                        = 19
	SYS_GETPID                       = 20
	SYS_MOUNT                        = 21
	SYS_UMOUNT                       = 22
	SYS_SETUID                       = 23
	SYS_GETUID                       = 24
	SYS_STIME                        = 25
	SYS_PTRACE                       = 26
	SYS_ALARM                        = 27
	SYS_OLDFSTAT                     = 28
	SYS_PAUSE                        = 29
	SYS_UTIME                        = 30
	SYS_STTY                         = 31
	SYS_GTTY                         = 32
	SYS_ACCESS                       = 33
	SYS_NICE                         = 34
	SYS_FTIME                        = 35
	SYS_SYNC                         = 36
	SYS_KILL                         = 37
	SYS_RENAME                       = 38
	SYS_MKDIR                        = 39
	SYS_RMDIR                        = 40
	SYS_DUP                          = 41
	SYS_PIPE                         = 42
	SYS_TIMES                        = 43
	SYS_PROF                         = 44
	SYS_BRK                          = 45
	SYS_SETGID                       = 46
	SYS_GETGID                       = 47
	SYS_SIGNAL                       = 48
	SYS_GETEUID                      = 49
	SYS_GETEGID                      = 50
	SYS_ACCT                         = 51
	SYS_UMOUNT2                      = 52
	SYS_LOCK                         = 53
	SYS_IOCTL                        = 54
	SYS_FCNTL                        = 55
	SYS_MPX                          = 56
	SYS_SETPGID                      = 57
	SYS_ULIMIT                       = 58
	SYS_OLDOLDUNAME                  = 59
	SYS_UMASK                        = 60
	SYS_CHROOT                       = 61
	SYS_USTAT                        = 62
	SYS_DUP2                         = 63
	SYS_GETPPID                      = 64
	SYS_GETPGRP                      = 65
	SYS_SETSID                       = 66
	SYS_SIGACTION                    = 67
	SYS_SGETMASK                     = 68
	SYS_SSETMASK                     = 69
	SYS_SETREUID                     = 70
	SYS_SETREGID                     = 71
	SYS_SIGSUSPEND                   = 72
	SYS_SIGPENDING                   = 73
	SYS_SETHOSTNAME                  = 74
	SYS_SETRLIMIT                    = 75
	SYS_GETRLIMIT                    = 76
	SYS_GETRUSAGE                    = 77
	SYS_GETTIMEOFDAY                 = 78
	SYS_SETTIMEOFDAY                 = 79
	SYS_GETGROUPS                    = 80
	SYS_SETGROUPS                    = 81
	SYS_SELECT                       = 82
	SYS_SYMLINK                      = 83
	SYS_OLDLSTAT                     = 84
	SYS_READLINK                     = 85
	SYS_USELIB                       = 86
	SYS_SWAPON                       = 87
	SYS_REBOOT                       = 88
	SYS_READDIR                      = 89
	SYS_MMAP                         = 90
	SYS_MUNMAP                       = 91
	SYS_TRUNCATE                     = 92
	SYS_FTRUNCATE                    = 93
	SYS_FCHMOD                       = 94
	SYS_FCHOWN                       = 95
	SYS_GETPRIORITY                  = 96
	SYS_SETPRIORITY                  = 97
	SYS_PROFIL                       = 98
	SYS_STATFS                       = 99
	SYS_FSTATFS                      = 100
	SYS_IOPERM                       = 101
	SYS_SOCKETCALL                   = 102
	SYS_SYSLOG                       = 103
	SYS_SETITIMER                    = 104
	SYS_GETITIMER                    = 105
	SYS_STAT                         = 106
	SYS_LSTAT                        = 107
	SYS_FSTAT                        = 108
	SYS_OLDUNAME                     = 109
	SYS_IOPL                         = 110
	SYS_VHANGUP                      = 111
	SYS_IDLE                         = 112
	SYS_VM86OLD                      = 113
	SYS_WAIT4                        = 114
	SYS_SWAPOFF                      = 115
	SYS_SYSINFO                      = 116
	SYS_IPC                          = 117
	SYS_FSYNC                        = 118
	SYS_SIGRETURN                    = 119
	SYS_CLONE                        = 120
	SYS_SETDOMAINNAME                = 121
	SYS_UNAME                        = 122
	SYS_MODIFY_LDT                   = 123
	SYS_ADJTIMEX                     = 124
	SYS_MPROTECT                     = 125
	SYS_SIGPROCMASK                  = 126
	SYS_CREATE_MODULE                = 127
	SYS_INIT_MODULE                  = 128
	SYS_DELETE_MODULE                = 129
	SYS_GET_KERNEL_SYMS              = 130
	SYS_QUOTACTL                     = 131
	SYS_GETPGID                      = 132
	SYS_FCHDIR                       = 133
	SYS_BDFLUSH                      = 134
	SYS_SYSFS                        = 135
	SYS_PERSONALITY                  = 136
	SYS_AFS_SYSCALL                  = 137
	SYS_SETFSUID                     = 138
	SYS_SETFSGID                     = 139
	SYS__LLSEEK                      = 140
	SYS_GETDENTS                     = 141
	SYS__NEWSELECT                   = 142
	SYS_FLOCK                        = 143
	SYS_MSYNC                        = 144
	SYS_READV                        = 145
	SYS_WRITEV                       = 146
	SYS_GETSID                       = 147
	SYS_FDATASYNC                    = 148
	SYS__SYSCTL                      = 149
	SYS_MLOCK                        = 150
	SYS_MUNLOCK                      = 151
	SYS_MLOCKALL                     = 152
	SYS_MUNLOCKALL                   = 153
	SYS_SCHED_SETPARAM               = 154
	SYS_SCHED_GETPARAM               = 155
	SYS_SCHED_SETSCHEDULER           = 156
	SYS_SCHED_GETSCHEDULER           = 157
	SYS_SCHED_YIELD                  = 158
	SYS_SCHED_GET_PRIORITY_MAX       = 159
	SYS_SCHED_GET_PRIORITY_MIN       = 160
	SYS_SCHED_RR_GET_INTERVAL        = 161
	SYS_NANOSLEEP                    = 162
	SYS_MREMAP                       = 163
	SYS_SETRESUID                    = 164
	SYS_GETRESUID                    = 165
	SYS_VM86                         = 166
	SYS_QUERY_MODULE                 = 167
	SYS_POLL                         = 168
	SYS_NFSSERVCTL                   = 169
	SYS_SETRESGID                    = 170
	SYS_GETRESGID                    = 171
	SYS_PRCTL                        = 172
	SYS_RT_SIGRETURN                 = 173
	SYS_RT_SIGACTION                 = 174
	SYS_RT_SIGPROCMASK               = 175
	SYS_RT_SIGPENDING                = 176
	SYS_RT_SIGTIMEDWAIT              = 177
	SYS_RT_SIGQUEUEINFO              = 178
	SYS_RT_SIGSUSPEND                = 179
	SYS_PREAD64                      = 180
	SYS_PWRITE64                     = 181
	SYS_CHOWN                        = 182
	SYS_GETCWD                       = 183
	SYS_CAPGET                       = 184
	SYS_CAPSET                       = 185
	SYS_SIGALTSTACK                  = 186
	SYS_SENDFILE                     = 187
	SYS_GETPMSG                      = 188
	SYS_PUTPMSG                      = 189
	SYS_VFORK                        = 190
	SYS_UGETRLIMIT                   = 191
	SYS_MMAP2                        = 192
	SYS_TRUNCATE64                   = 193
	SYS_FTRUNCATE64                  = 194
	SYS_STAT64                       = 195
	SYS_LSTAT64                      = 196
	SYS_FSTAT64                      = 197
	SYS_LCHOWN32                     = 198
	SYS_GETUID32                     = 199
	SYS_GETGID32                     = 200
	SYS_GETEUID32                    = 201
	SYS_GETEGID32                    = 202
	SYS_SETREUID32                   = 203
	SYS_SETREGID32                   = 204
	SYS_GETGROUPS32                  = 205
	SYS_SETGROUPS32                  = 206
	SYS_FCHOWN32                     = 207
	SYS_SETRESUID32                  = 208
	SYS_GETRESUID32                  = 209
	SYS_SETRESGID32                  = 210
	SYS_GETRESGID32                  = 211
	SYS_CHOWN32                      = 212
	SYS_SETUID32                     = 213
	SYS_SETGID32                     = 214
	SYS_SETFSUID32                   = 215
	SYS_SETFSGID32                   = 216
	SYS_PIVOT_ROOT                   = 217
	SYS_MINCORE                      = 218
	SYS_MADVISE                      = 219
	SYS_GETDENTS64                   = 220
	SYS_FCNTL64                      = 221
	SYS_GETTID                       = 224
	SYS_READAHEAD                    = 225
	SYS_SETXATTR                     = 226
	SYS_LSETXATTR                    = 227
	SYS_FSETXATTR                    = 228
	SYS_GETXATTR                     = 229
	SYS_LGETXATTR                    = 230
	SYS_FGETXATTR                    = 231
	SYS_LISTXATTR                    = 232
	SYS_LLISTXATTR                   = 233
	SYS_FLISTXATTR                   = 234
	SYS_REMOVEXATTR                  = 235
	SYS_LREMOVEXATTR                 = 236
	SYS_FREMOVEXATTR                 = 237
	SYS_TKILL                        = 238
	SYS_SENDFILE64                   = 239
	SYS_FUTEX                        = 240
	SYS_SCHED_SETAFFINITY            = 241
	SYS_SCHED_GETAFFINITY            = 242
	SYS_SET_THREAD_AREA              = 243
	SYS_GET_THREAD_AREA              = 244
	SYS_IO_SETUP                     = 245
	SYS_IO_DESTROY                   = 246
	SYS_IO_GETEVENTS                 = 247
	SYS_IO_SUBMIT                    = 248
	SYS_IO_CANCEL                    = 249
	SYS_FADVISE64                    = 250
	SYS_EXIT_GROUP                   = 252
	SYS_LOOKUP_DCOOKIE               = 253
	SYS_EPOLL_CREATE                 = 254
	SYS_EPOLL_CTL                    = 255
	SYS_EPOLL_WAIT                   = 256
	SYS_REMAP_FILE_PAGES             = 257
	SYS_SET_TID_ADDRESS              = 258
	SYS_TIMER_CREATE                 = 259
	SYS_TIMER_SETTIME                = 260
	SYS_TIMER_GETTIME                = 261
	SYS_TIMER_GETOVERRUN             = 262
	SYS_TIMER_DELETE                 = 263
	SYS_CLOCK_SETTIME                = 264
	SYS_CLOCK_GETTIME                = 265
	SYS_CLOCK_GETRES                 = 266
	SYS_CLOCK_NANOSLEEP              = 267
	SYS_STATFS64                     = 268
	SYS_FSTATFS64                    = 269
	SYS_TGKILL                       = 270
	SYS_UTIMES                       = 271
	SYS_FADVISE64_64                 = 272
	SYS_VSERVER                      = 273
	SYS_MBIND                        = 274
	SYS_GET_MEMPOLICY                = 275
	SYS_SET_MEMPOLICY                = 276
	SYS_MQ_OPEN                      = 277
	SYS_MQ_UNLINK                    = 278
	SYS_MQ_TIMEDSEND                 = 279
	SYS_MQ_TIMEDRECEIVE              = 280
	SYS_MQ_NOTIFY                    = 281
	SYS_MQ_GETSETATTR                = 282
	SYS_KEXEC_LOAD                   = 283
	SYS_WAITID                       = 284
	SYS_ADD_KEY                      = 286
	SYS_REQUEST_KEY                  = 287
	SYS_KEYCTL                       = 288
	SYS_IOPRIO_SET                   = 289
	SYS_IOPRIO_GET                   = 290
	SYS_INOTIFY_INIT                 = 291
	SYS_INOTIFY_ADD_WATCH            = 292
	SYS_INOTIFY_RM_WATCH             = 293
	SYS_MIGRATE_PAGES                = 294
	SYS_OPENAT                       = 295
	SYS_MKDIRAT                      = 296
	SYS_MKNODAT                      = 297
	SYS_FCHOWNAT                     = 298
	SYS_FUTIMESAT                    = 299
	SYS_FSTATAT64                    = 300
	SYS_UNLINKAT                     = 301
	SYS_RENAMEAT                     = 302
	SYS_LINKAT                       = 303
	SYS_SYMLINKAT                    = 304
	SYS_READLINKAT                   = 305
	SYS_FCHMODAT                     = 306
	SYS_FACCESSAT                    = 307
	SYS_PSELECT6                     = 308
	SYS_PPOLL                        = 309
	SYS_UNSHARE                      = 310
	SYS_SET_ROBUST_LIST              = 311
	SYS_GET_ROBUST_LIST              = 312
	SYS_SPLICE                       = 313
	SYS_SYNC_FILE_RANGE              = 314
	SYS_TEE                          = 315
	SYS_VMSPLICE                     = 316
	SYS_MOVE_PAGES                   = 317
	SYS_GETCPU                       = 318
	SYS_EPOLL_PWAIT                  = 319
	SYS_UTIMENSAT                    = 320
	SYS_SIGNALFD                     = 321
	SYS_TIMERFD_CREATE               = 322
	SYS_EVENTFD                      = 323
	SYS_FALLOCATE                    = 324
	SYS_TIMERFD_SETTIME              = 325
	SYS_TIMERFD_GETTIME              = 326
	SYS_SIGNALFD4                    = 327
	SYS_EVENTFD2                     = 328
	SYS_EPOLL_CREATE1                = 329
	SYS_DUP3                         = 330
	SYS_PIPE2                        = 331
	SYS_INOTIFY_INIT1                = 332
	SYS_PREADV                       = 333
	SYS_PWRITEV                      = 334
	SYS_RT_TGSIGQUEUEINFO            = 335
	SYS_PERF_EVENT_OPEN              = 336
	SYS_RECVMMSG                     = 337
	SYS_FANOTIFY_INIT                = 338
	SYS_FANOTIFY_MARK                = 339
	SYS_PRLIMIT64                    = 340
	SYS_NAME_TO_HANDLE_AT            = 341
	SYS_OPEN_BY_HANDLE_AT            = 342
	SYS_CLOCK_ADJTIME                = 343
	SYS_SYNCFS                       = 344
	SYS_SENDMMSG                     = 345
	SYS_SETNS                        = 346
	SYS_PROCESS_VM_READV             = 347
	SYS_PROCESS_VM_WRITEV            = 348
	SYS_KCMP                         = 349
	SYS_FINIT_MODULE                 = 350
	SYS_SCHED_SETATTR                = 351
	SYS_SCHED_GETATTR                = 352
	SYS_RENAMEAT2                    = 353
	SYS_SECCOMP                      = 354
	SYS_GETRANDOM                    = 355
	SYS_MEMFD_CREATE                 = 356
	SYS_BPF                          = 357
	SYS_EXECVEAT                     = 358
	SYS_SOCKET                       = 359
	SYS_SOCKETPAIR                   = 360
	SYS_BIND                         = 361
	SYS_CONNECT                      = 362
	SYS_LISTEN                       = 363
	SYS_ACCEPT4                      = 364
	SYS_GETSOCKOPT                   = 365
	SYS_SETSOCKOPT                   = 366
	SYS_GETSOCKNAME                  = 367
	SYS_GETPEERNAME                  = 368
	SYS_SENDTO                       = 369
	SYS_SENDMSG                      = 370
	SYS_RECVFROM                     = 371
	SYS_RECVMSG                      = 372
	SYS_SHUTDOWN                     = 373
	SYS_USERFAULTFD                  = 374
	SYS_MEMBARRIER                   = 375
	SYS_MLOCK2                       = 376
	SYS_COPY_FILE_RANGE              = 377
	SYS_PREADV2                      = 378
	SYS_PWRITEV2                     = 379
	SYS_PKEY_MPROTECT                = 380
	SYS_PKEY_ALLOC                   = 381
	SYS_PKEY_FREE                    = 382
	SYS_STATX                        = 383
	SYS_ARCH_PRCTL                   = 384
	SYS_IO_PGETEVENTS                = 385
	SYS_RSEQ                         = 386
	SYS_SEMGET                       = 393
	SYS_SEMCTL                       = 394
	SYS_SHMGET                       = 395
	SYS_SHMCTL                       = 396
	SYS_SHMAT                        = 397
	SYS_SHMDT                        = 398
	SYS_MSGGET                       = 399
	SYS_MSGSND                       = 400
	SYS_MSGRCV                       = 401
	SYS_MSGCTL                       = 402
	SYS_CLOCK_GETTIME64              = 403
	SYS_CLOCK_SETTIME64              = 404
	SYS_CLOCK_ADJTIME64              = 405
	SYS_CLOCK_GETRES_TIME64          = 406
	SYS_CLOCK_NANOSLEEP_TIME64       = 407
	SYS_TIMER_GETTIME64              = 408
	SYS_TIMER_SETTIME64              = 409
	SYS_TIMERFD_GETTIME64            = 410
	SYS_TIMERFD_SETTIME64            = 411
	SYS_UTIMENSAT_TIME64             = 412
	SYS_PSELECT6_TIME64              = 413
	SYS_PPOLL_TIME64                 = 414
	SYS_IO_PGETEVENTS_TIME64         = 416
	SYS_RECVMMSG_TIME64              = 417
	SYS_MQ_TIMEDSEND_TIME64          = 418
	SYS_MQ_TIMEDRECEIVE_TIME64       = 419
	SYS_SEMTIMEDOP_TIME64            = 420
	SYS_RT_SIGTIMEDWAIT_TIME64       = 421
	SYS_FUTEX_TIME64                 = 422
	SYS_SCHED_RR_GET_INTERVAL_TIME64 = 423
	SYS_PIDFD_SEND_SIGNAL            = 424
	SYS_IO_URING_SETUP               = 425
	SYS_IO_URING_ENTER               = 426
	SYS_IO_URING_REGISTER            = 427
	SYS_OPEN_TREE                    = 428
	SYS_MOVE_MOUNT                   = 429
	SYS_FSOPEN                       = 430
	SYS_FSCONFIG                     = 431
	SYS_FSMOUNT                      = 432
	SYS_FSPICK                       = 433
	SYS_PIDFD_OPEN                   = 434
	SYS_CLONE3                       = 435
	SYS_CLOSE_RANGE                  = 436
	SYS_OPENAT2                      = 437
	SYS_PIDFD_GETFD                  = 438
	SYS_FACCESSAT2                   = 439
	SYS_PROCESS_MADVISE              = 440
	SYS_EPOLL_PWAIT2                 = 441
	SYS_MOUNT_SETATTR                = 442
	SYS_QUOTACTL_FD                  = 443
	SYS_LANDLOCK_CREATE_RULESET      = 444
	SYS_LANDLOCK_ADD_RULE            = 445
	SYS_LANDLOCK_RESTRICT_SELF       = 446
	SYS_MEMFD_SECRET                 = 447
	SYS_PROCESS_MRELEASE             = 448
	SYS_FUTEX_WAITV                  = 449
	SYS_SET_MEMPOLICY_HOME_NODE      = 450
	SYS_CACHESTAT                    = 451
	SYS_FCHMODAT2                    = 452
	SYS_MAP_SHADOW_STACK             = 453
	SYS_FUTEX_WAKE                   = 454
	SYS_FUTEX_WAIT                   = 455
	SYS_FUTEX_REQUEUE                = 456
	SYS_STATMOUNT                    = 457
	SYS_LISTMOUNT                    = 458
	SYS_LSM_GET_SELF_ATTR            = 459
	SYS_LSM_SET_SELF_ATTR            = 460
	SYS_LSM_LIST_MODULES             = 461
	SYS_MSEAL                        = 462
)

var Names = map[int]string{
	SYS_RESTART_SYSCALL:              "restart_syscall",
	SYS_EXIT:                         "exit",
	SYS_FORK:                         "fork",
	SYS_READ:                         "read",
	SYS_WRITE:                        "write",
	SYS_OPEN:                         "open",
	SYS_CLOSE:                        "close",
	SYS_WAITPID:                      "waitpid",
	SYS_CREAT:                        "creat",
	SYS_LINK:                         "link",
	SYS_UNLINK:                       "unlink",
	SYS_EXECVE:                       "execve",
	SYS_CHDIR:                        "chdir",
	SYS_TIME:                         "time",
	SYS_MKNOD:                        "mknod",
	SYS_CHMOD:                        "chmod",
	SYS_LCHOWN:                       "lchown",
	SYS_BREAK:                        "break",
	SYS_OLDSTAT:                      "oldstat",
	SYS_LSEEK:                        "lseek",
	SYS_GETPID:                       "getpid",
	SYS_MOUNT:                        "mount",
	SYS_UMOUNT:                       "umount",
	SYS_SETUID:                       "setuid",
	SYS_GETUID:                       "getuid",
	SYS_STIME:                        "stime",
	SYS_PTRACE:                       "ptrace",
	SYS_ALARM:                        "alarm",
	SYS_OLDFSTAT:                     "oldfstat",
	SYS_PAUSE:                        "pause",
	SYS_UTIME:                        "utime",
	SYS_STTY:                         "stty",
	SYS_GTTY:                         "gtty",
	SYS_ACCESS:                       "access",
	SYS_NICE:                         "nice",
	SYS_FTIME:                        "ftime",
	SYS_SYNC:                         "sync",
	SYS_KILL:                         "kill",
	SYS_RENAME:                       "rename",
	SYS_MKDIR:                        "mkdir",
	SYS_RMDIR:                        "rmdir",
	SYS_DUP:                          "dup",
	SYS_PIPE:                         "pipe",
	SYS_TIMES:                        "times",
	SYS_PROF:                         "prof",
	SYS_BRK:                          "brk",
	SYS_SETGID:                       "setgid",
	SYS_GETGID:                       "getgid",
	SYS_SIGNAL:                       "signal",
	SYS_GETEUID:                      "geteuid",
	SYS_GETEGID:                      "getegid",
	SYS_ACCT:                         "acct",
	SYS_UMOUNT2:                      "umount2",
	SYS_LOCK:                         "lock",
	SYS_IOCTL:                        "ioctl",
	SYS_FCNTL:                        "fcntl",
	SYS_MPX:                          "mpx",
	SYS_SETPGID:                      "setpgid",
	SYS_ULIMIT:                       "ulimit",
	SYS_OLDOLDUNAME:                  "oldolduname",
	SYS_UMASK:                        "umask",
	SYS_CHROOT:                       "chroot",
	SYS_USTAT:                        "ustat",
	SYS_DUP2:                         "dup2",
	SYS_GETPPID:                      "getppid",
	SYS_GETPGRP:                      "getpgrp",
	SYS_SETSID:                       "setsid",
	SYS_SIGACTION:                    "sigaction",
	SYS_SGETMASK:                     "sgetmask",
	SYS_SSETMASK:                     "ssetmask",
	SYS_SETREUID:                     "setreuid",
	SYS_SETREGID:                     "setregid",
	SYS_SIGSUSPEND:                   "sigsuspend",
	SYS_SIGPENDING:                   "sigpending",
	SYS_SETHOSTNAME:                  "sethostname",
	SYS_SETRLIMIT:                    "setrlimit",
	SYS_GETRLIMIT:                    "getrlimit",
	SYS_GETRUSAGE:                    "getrusage",
	SYS_GETTIMEOFDAY:                 "gettimeofday",
	SYS_SETTIMEOFDAY:                 "settimeofday",
	SYS_GETGROUPS:                    "getgroups",
	SYS_SETGROUPS:                    "setgroups",
	SYS_SELECT:                       "select",
	SYS_SYMLINK:                      "symlink",
	SYS_OLDLSTAT:                     "oldlstat",
	SYS_READLINK:                     "readlink",
	SYS_USELIB:                       "uselib",
	SYS_SWAPON:                       "swapon",
	SYS_REBOOT:                       "reboot",
	SYS_READDIR:                      "readdir",
	SYS_MMAP:                         "mmap",
	SYS_MUNMAP:                       "munmap",
	SYS_TRUNCATE:                     "truncate",
	SYS_FTRUNCATE:                    "ftruncate",
	SYS_FCHMOD:                       "fchmod",
	SYS_FCHOWN:                       "fchown",
	SYS_GETPRIORITY:                  "getpriority",
	SYS_SETPRIORITY:                  "setpriority",
	SYS_PROFIL:                       "profil",
	SYS_STATFS:                       "statfs",
	SYS_FSTATFS:                      "fstatfs",
	SYS_IOPERM:                       "ioperm",
	SYS_SOCKETCALL:                   "socketcall",
	SYS_SYSLOG:                       "syslog",
	SYS_SETITIMER:                    "setitimer",
	SYS_GETITIMER:                    "getitimer",
	SYS_STAT:                         "stat",
	SYS_LSTAT:                        "lstat",
	SYS_FSTAT:                        "fstat",
	SYS_OLDUNAME:                     "olduname",
	SYS_IOPL:                         "iopl",
	SYS_VHANGUP:                      "vhangup",
	SYS_IDLE:                         "idle",
	SYS_VM86OLD:                      "vm86old",
	SYS_WAIT4:                        "wait4",
	SYS_SWAPOFF:                      "swapoff",
	SYS_SYSINFO:                      "sysinfo",
	SYS_IPC:                          "ipc",
	SYS_FSYNC:                        "fsync",
	SYS_SIGRETURN:                    "sigreturn",
	SYS_CLONE:                        "clone",
	SYS_SETDOMAINNAME:                "setdomainname",
	SYS_UNAME:                        "uname",
	SYS_MODIFY_LDT:                   "modify_ldt",
	SYS_ADJTIMEX:                     "adjtimex",
	SYS_MPROTECT:                     "mprotect",
	SYS_SIGPROCMASK:                  "sigprocmask",
	SYS_CREATE_MODULE:                "create_module",
	SYS_INIT_MODULE:                  "init_module",
	SYS_DELETE_MODULE:                "delete_module",
	SYS_GET_KERNEL_SYMS:              "get_kernel_syms",
	SYS_QUOTACTL:                     "quotactl",
	SYS_GETPGID:                      "getpgid",
	SYS_FCHDIR:                       "fchdir",
	SYS_BDFLUSH:                      "bdflush",
	SYS_SYSFS:                        "sysfs",
	SYS_PERSONALITY:                  "personality",
	SYS_AFS_SYSCALL:                  "afs_syscall",
	SYS_SETFSUID:                     "setfsuid",
	SYS_SETFSGID:                     "setfsgid",
	SYS__LLSEEK:                      "_llseek",
	SYS_GETDENTS:                     "getdents",
	SYS__NEWSELECT:                   "_newselect",
	SYS_FLOCK:                        "flock",
	SYS_MSYNC:                        "msync",
	SYS_READV:                        "readv",
	SYS_WRITEV:                       "writev",
	SYS_GETSID:                       "getsid",
	SYS_FDATASYNC:                    "fdatasync",
	SYS__SYSCTL:                      "_sysctl",
	SYS_MLOCK:                        "mlock",
	SYS_MUNLOCK:                      "munlock",
	SYS_MLOCKALL:                     "mlockall",
	SYS_MUNLOCKALL:                   "munlockall",
	SYS_SCHED_SETPARAM:               "sched_setparam",
	SYS_SCHED_GETPARAM:               "sched_getparam",
	SYS_SCHED_SETSCHEDULER:           "sched_setscheduler",
	SYS_SCHED_GETSCHEDULER:           "sched_getscheduler",
	SYS_SCHED_YIELD:                  "sched_yield",
	SYS_SCHED_GET_PRIORITY_MAX:       "sched_get_priority_max",
	SYS_SCHED_GET_PRIORITY_MIN:       "sched_get_priority_min",
	SYS_SCHED_RR_GET_INTERVAL:        "sched_rr_get_interval",
	SYS_NANOSLEEP:                    "nanosleep",
	SYS_MREMAP:                       "mremap",
	SYS_SETRESUID:                    "setresuid",
	SYS_GETRESUID:                    "getresuid",
	SYS_VM86:                         "vm86",
	SYS_QUERY_MODULE:                 "query_module",
	SYS_POLL:                         "poll",
	SYS_NFSSERVCTL:                   "nfsservctl",
	SYS_SETRESGID:                    "setresgid",
	SYS_GETRESGID:                    "getresgid",
	SYS_PRCTL:                        "prctl",
	SYS_RT_SIGRETURN:                 "rt_sigreturn",
	SYS_RT_SIGACTION:                 "rt_sigaction",
	SYS_RT_SIGPROCMASK:               "rt_sigprocmask",
	SYS_RT_SIGPENDING:                "rt_sigpending",
	SYS_RT_SIGTIMEDWAIT:              "rt_sigtimedwait",
	SYS_RT_SIGQUEUEINFO:              "rt_sigqueueinfo",
	SYS_RT_SIGSUSPEND:                "rt_sigsuspend",
	SYS_PREAD64:                      "pread64",
	SYS_PWRITE64:                     "pwrite64",
	SYS_CHOWN:                        "chown",
	SYS_GETCWD:                       "getcwd",
	SYS_CAPGET:                       "capget",
	SYS_CAPSET:                       "capset",
	SYS_SIGALTSTACK:                  "sigaltstack",
	SYS_SENDFILE:                     "sendfile",
	SYS_GETPMSG:                      "getpmsg",
	SYS_PUTPMSG:                      "putpmsg",
	SYS_VFORK:                        "vfork",
	SYS_UGETRLIMIT:                   "ugetrlimit",
	SYS_MMAP2:                        "mmap2",
	SYS_TRUNCATE64:                   "truncate64",
	SYS_FTRUNCATE64:                  "ftruncate64",
	SYS_STAT64:                       "stat64",
	SYS_LSTAT64:                      "lstat64",
	SYS_FSTAT64:                      "fstat64",
	SYS_LCHOWN32:                     "lchown32",
	SYS_GETUID32:                     "getuid32",
	SYS_GETGID32:                     "getgid32",
	SYS_GETEUID32:                    "geteuid32",
	SYS_GETEGID32:                    "getegid32",
	SYS_SETREUID32:                   "setreuid32",
	SYS_SETREGID32:                   "setregid32",
	SYS_GETGROUPS32:                  "getgroups32",
	SYS_SETGROUPS32:                  "setgroups32",
	SYS_FCHOWN32:                     "fchown32",
	SYS_SETRESUID32:                  "setresuid32",
	SYS_GETRESUID32:                  "getresuid32",
	SYS_SETRESGID32:                  "setresgid32",
	SYS_GETRESGID32:                  "getresgid32",
	SYS_CHOWN32:                      "chown32",
	SYS_SETUID32:                     "setuid32",
	SYS_SETGID32:                     "setgid32",
	SYS_SETFSUID32:                   "setfsuid32",
	SYS_SETFSGID32:                   "setfsgid32",
	SYS_PIVOT_ROOT:                   "pivot_root",
	SYS_MINCORE:                      "mincore",
	SYS_MADVISE:                      "madvise",
	SYS_GETDENTS64:                   "getdents64",
	SYS_FCNTL64:                      "fcntl64",
	SYS_GETTID:                       "gettid",
	SYS_READAHEAD:                    "readahead",
	SYS_SETXATTR:                     "setxattr",
	SYS_LSETXATTR:                    "lsetxattr",
	SYS_FSETXATTR:                    "fsetxattr",
	SYS_GETXATTR:                     "getxattr",
	SYS_LGETXATTR:                    "lgetxattr",
	SYS_FGETXATTR:                    "fgetxattr",
	SYS_LISTXATTR:                    "listxattr",
	SYS_LLISTXATTR:                   "llistxattr",
	SYS_FLISTXATTR:                   "flistxattr",
	SYS_REMOVEXATTR:                  "removexattr",
	SYS_LREMOVEXATTR:                 "lremovexattr",
	SYS_FREMOVEXATTR:                 "fremovexattr",
	SYS_TKILL:                        "tkill",
	SYS_SENDFILE64:                   "sendfile64",
	SYS_FUTEX:                        "futex",
	SYS_SCHED_SETAFFINITY:            "sched_setaffinity",
	SYS_SCHED_GETAFFINITY:            "sched_getaffinity",
	SYS_SET_THREAD_AREA:              "set_thread_area",
	SYS_GET_THREAD_AREA:              "get_thread_area",
	SYS_IO_SETUP:                     "io_setup",
	SYS_IO_DESTROY:                   "io_destroy",
	SYS_IO_GETEVENTS:                 "io_getevents",
	SYS_IO_SUBMIT:                    "io_submit",
	SYS_IO_CANCEL:                    "io_cancel",
	SYS_FADVISE64:                    "fadvise64",
	SYS_EXIT_GROUP:                   "exit_group",
	SYS_LOOKUP_DCOOKIE:               "lookup_dcookie",
	SYS_EPOLL_CREATE:                 "epoll_create",
	SYS_EPOLL_CTL:                    "epoll_ctl",
	SYS_EPOLL_WAIT:                   "epoll_wait",
	SYS_REMAP_FILE_PAGES:             "remap_file_pages",
	SYS_SET_TID_ADDRESS:              "set_tid_address",
	SYS_TIMER_CREATE:                 "timer_create",
	SYS_TIMER_SETTIME:                "timer_settime",
	SYS_TIMER_GETTIME:                "timer_gettime",
	SYS_TIMER_GETOVERRUN:             "timer_getoverrun",
	SYS_TIMER_DELETE:                 "timer_delete",
	SYS_CLOCK_SETTIME:                "clock_settime",
	SYS_CLOCK_GETTIME:                "clock_gettime",
	SYS_CLOCK_GETRES:                 "clock_getres",
	SYS_CLOCK_NANOSLEEP:              "clock_nanosleep",
	SYS_STATFS64:                     "statfs64",
	SYS_FSTATFS64:                    "fstatfs64",
	SYS_TGKILL:                       "tgkill",
	SYS_UTIMES:                       "utimes",
	SYS_FADVISE64_64:                 "fadvise64_64",
	SYS_VSERVER:                      "vserver",
	SYS_MBIND:                        "mbind",
	SYS_GET_MEMPOLICY:                "get_mempolicy",
	SYS_SET_MEMPOLICY:                "set_mempolicy",
	SYS_MQ_OPEN:                      "mq_open",
	SYS_MQ_UNLINK:                    "mq_unlink",
	SYS_MQ_TIMEDSEND:                 "mq_timedsend",
	SYS_MQ_TIMEDRECEIVE:              "mq_timedreceive",
	SYS_MQ_NOTIFY:                    "mq_notify",
	SYS_MQ_GETSETATTR:                "mq_getsetattr",
	SYS_KEXEC_LOAD:                   "kexec_load",
	SYS_WAITID:                       "waitid",
	SYS_ADD_KEY:                      "add_key",
	SYS_REQUEST_KEY:                  "request_key",
	SYS_KEYCTL:                       "keyctl",
	SYS_IOPRIO_SET:                   "ioprio_set",
	SYS_IOPRIO_GET:                   "ioprio_get",
	SYS_INOTIFY_INIT:                 "inotify_init",
	SYS_INOTIFY_ADD_WATCH:            "inotify_add_watch",
	SYS_INOTIFY_RM_WATCH:             "inotify_rm_watch",
	SYS_MIGRATE_PAGES:                "migrate_pages",
	SYS_OPENAT:                       "openat",
	SYS_MKDIRAT:                      "mkdirat",
	SYS_MKNODAT:                      "mknodat",
	SYS_FCHOWNAT:                     "fchownat",
	SYS_FUTIMESAT:                    "futimesat",
	SYS_FSTATAT64:                    "fstatat64",
	SYS_UNLINKAT:                     "unlinkat",
	SYS_RENAMEAT:                     "renameat",
	SYS_LINKAT:                       "linkat",
	SYS_SYMLINKAT:                    "symlinkat",
	SYS_READLINKAT:                   "readlinkat",
	SYS_FCHMODAT:                     "fchmodat",
	SYS_FACCESSAT:                    "faccessat",
	SYS_PSELECT6:                     "pselect6",
	SYS_PPOLL:                        "ppoll",
	SYS_UNSHARE:                      "unshare",
	SYS_SET_ROBUST_LIST:              "set_robust_list",
	SYS_GET_ROBUST_LIST:              "get_robust_list",
	SYS_SPLICE:                       "splice",
	SYS_SYNC_FILE_RANGE:              "sync_file_range",
	SYS_TEE:                          "tee",
	SYS_VMSPLICE:                     "vmsplice",
	SYS_MOVE_PAGES:                   "move_pages",
	SYS_GETCPU:                       "getcpu",
	SYS_EPOLL_PWAIT:                  "epoll_pwait",
	SYS_UTIMENSAT:                    "utimensat",
	SYS_SIGNALFD:                     "signalfd",
	SYS_TIMERFD_CREATE:               "timerfd_create",
	SYS_EVENTFD:                      "eventfd",
	SYS_FALLOCATE:                    "fallocate",
	SYS_TIMERFD_SETTIME:              "timerfd_settime",
	SYS_TIMERFD_GETTIME:              "timerfd_gettime",
	SYS_SIGNALFD4:                    "signalfd4",
	SYS_EVENTFD2:                     "eventfd2",
	SYS_EPOLL_CREATE1:                "epoll_create1",
	SYS_DUP3:                         "dup3",
	SYS_PIPE2:                        "pipe2",
	SYS_INOTIFY_INIT1:                "inotify_init1",
	SYS_PREADV:                       "preadv",
	SYS_PWRITEV:                      "pwritev",
	SYS_RT_TGSIGQUEUEINFO:            "rt_tgsigqueueinfo",
	SYS_PERF_EVENT_OPEN:              "perf_event_open",
	SYS_RECVMMSG:                     "recvmmsg",
	SYS_FANOTIFY_INIT:                "fanotify_init",
	SYS_FANOTIFY_MARK:                "fanotify_mark",
	SYS_PRLIMIT64:                    "prlimit64",
	SYS_NAME_TO_HANDLE_AT:            "name_to_handle_at",
	SYS_OPEN_BY_HANDLE_AT:            "open_by_handle_at",
	SYS_CLOCK_ADJTIME:                "clock_adjtime",
	SYS_SYNCFS:                       "syncfs",
	SYS_SENDMMSG:                     "sendmmsg",
	SYS_SETNS:                        "setns",
	SYS_PROCESS_VM_READV:             "process_vm_readv",
	SYS_PROCESS_VM_WRITEV:            "process_vm_writev",
	SYS_KCMP:                         "kcmp",
	SYS_FINIT_MODULE:                 "finit_module",
	SYS_SCHED_SETATTR:                "sched_setattr",
	SYS_SCHED_GETATTR:                "sched_getattr",
	SYS_RENAMEAT2:                    "renameat2",
	SYS_SECCOMP:                      "seccomp",
	SYS_GETRANDOM:                    "getrandom",
	SYS_MEMFD_CREATE:                 "memfd_create",
	SYS_BPF:                          "bpf",
	SYS_EXECVEAT:                     "execveat",
	SYS_SOCKET:                       "socket",
	SYS_SOCKETPAIR:                   "socketpair",
	SYS_BIND:                         "bind",
	SYS_CONNECT:                      "connect",
	SYS_LISTEN:                       "listen",
	SYS_ACCEPT4:                      "accept4",
	SYS_GETSOCKOPT:                   "getsockopt",
	SYS_SETSOCKOPT:                   "setsockopt",
	SYS_GETSOCKNAME:                  "getsockname",
	SYS_GETPEERNAME:                  "getpeername",
	SYS_SENDTO:                       "sendto",
	SYS_SENDMSG:                      "sendmsg",
	SYS_RECVFROM:                     "recvfrom",
	SYS_RECVMSG:                      "recvmsg",
	SYS_SHUTDOWN:                     "shutdown",
	SYS_USERFAULTFD:                  "userfaultfd",
	SYS_MEMBARRIER:                   "membarrier",
	SYS_MLOCK2:                       "mlock2",
	SYS_COPY_FILE_RANGE:              "copy_file_range",
	SYS_PREADV2:                      "preadv2",
	SYS_PWRITEV2:                     "pwritev2",
	SYS_PKEY_MPROTECT:                "pkey_mprotect",
	SYS_PKEY_ALLOC:                   "pkey_alloc",
	SYS_PKEY_FREE:                    "pkey_free",
	SYS_STATX:                        "statx",
	SYS_ARCH_PRCTL:                   "arch_prctl",
	SYS_IO_PGETEVENTS:                "io_pgetevents",
	SYS_RSEQ:                         "rseq",
	SYS_SEMGET:                       "semget",
	SYS_SEMCTL:                       "semctl",
	SYS_SHMGET:                       "shmget",
	SYS_SHMCTL:                       "shmctl",
	SYS_SHMAT:                        "shmat",
	SYS_SHMDT:                        "shmdt",
	SYS_MSGGET:                       "msgget",
	SYS_MSGSND:                       "msgsnd",
	SYS_MSGRCV:                       "msgrcv",
	SYS_MSGCTL:                       "msgctl",
	SYS_CLOCK_GETTIME64:              "clock_gettime64",
	SYS_CLOCK_SETTIME64:              "clock_settime64",
	SYS_CLOCK_ADJTIME64:              "clock_adjtime64",
	SYS_CLOCK_GETRES_TIME64:          "clock_getres_time64",
	SYS_CLOCK_NANOSLEEP_TIME64:       "clock_nanosleep_time64",
	SYS_TIMER_GETTIME64:              "timer_gettime64",
	SYS_TIMER_SETTIME64:              "timer_settime64",
	SYS_TIMERFD_GETTIME64:            "timerfd_gettime64",
	SYS_TIMERFD_SETTIME64:            "timerfd_settime64",
	SYS_UTIMENSAT_TIME64:             "utimensat_time64",
	SYS_PSELECT6_TIME64:              "pselect6_time64",
	SYS_PPOLL_TIME64:                 "ppoll_time64",
	SYS_IO_PGETEVENTS_TIME64:         "io_pgetevents_time64",
	SYS_RECVMMSG_TIME64:              "recvmmsg_time64",
	SYS_MQ_TIMEDSEND_TIME64:          "mq_timedsend_time64",
	SYS_MQ_TIMEDRECEIVE_TIME64:       "mq_timedreceive_time64",
	SYS_SEMTIMEDOP_TIME64:            "semtimedop_time64",
	SYS_RT_SIGTIMEDWAIT_TIME64:       "rt_sigtimedwait_time64",
	SYS_FUTEX_TIME64:                 "futex_time64",
	SYS_SCHED_RR_GET_INTERVAL_TIME64: "sched_rr_get_interval_time64",
	SYS_PIDFD_SEND_SIGNAL:            "pidfd_send_signal",
	SYS_IO_URING_SETUP:               "io_uring_setup",
	SYS_IO_URING_ENTER:               "io_uring_enter",
	SYS_IO_URING_REGISTER:            "io_uring_register",
	SYS_OPEN_TREE:                    "open_tree",
	SYS_MOVE_MOUNT:                   "move_mount",
	SYS_FSOPEN:                       "fsopen",
	SYS_FSCONFIG:                     "fsconfig",
	SYS_FSMOUNT:                      "fsmount",
	SYS_FSPICK:                       "fspick",
	SYS_PIDFD_OPEN:                   "pidfd_open",
	SYS_CLONE3:                       "clone3",
	SYS_CLOSE_RANGE:                  "close_range",
	SYS_OPENAT2:                      "openat2",
	SYS_PIDFD_GETFD:                  "pidfd_getfd",
	SYS_FACCESSAT2:                   "faccessat2",
	SYS_PROCESS_MADVISE:              "process_madvise",
	SYS_EPOLL_PWAIT2:                 "epoll_pwait2",
	SYS_MOUNT_SETATTR:                "mount_setattr",
	SYS_QUOTACTL_FD:                  "quotactl_fd",
	SYS_LANDLOCK_CREATE_RULESET:      "landlock_create_ruleset",
	SYS_LANDLOCK_ADD_RULE:            "landlock_add_rule",
	SYS_LANDLOCK_RESTRICT_SELF:       "landlock_restrict_self",
	SYS_MEMFD_SECRET:                 "memfd_secret",
	SYS_PROCESS_MRELEASE:             "process_mrelease",
	SYS_FUTEX_WAITV:                  "futex_waitv",
	SYS_SET_MEMPOLICY_HOME_NODE:      "set_mempolicy_home_node",
	SYS_CACHESTAT:                    "cachestat",
	SYS_FCHMODAT2:                    "fchmodat2",
	SYS_MAP_SHADOW_STACK:             "map_shadow_stack",
	SYS_FUTEX_WAKE:                   "futex_wake",
	SYS_FUTEX_WAIT:                   "futex_wait",
	SYS_FUTEX_REQUEUE:                "futex_requeue",
	SYS_STATMOUNT:                    "statmount",
	SYS_LISTMOUNT:                    "listmount",
	SYS_LSM_GET_SELF_ATTR:            "lsm_get_self_attr",
	SYS_LSM_SET_SELF_ATTR:            "lsm_set_self_attr",
	SYS_LSM_LIST_MODULES:             "lsm_list_modules",
	SYS_MSEAL:                        "mseal",
}
