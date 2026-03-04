#ifndef WIREGUARD_H
#define WIREGUARD_H

#include <sys/types.h>
#include <stdint.h>
#include <stdbool.h>

typedef void(*logger_fn_t)(void *context, int level, const char *msg);
extern void pltSetLogger(void *context, logger_fn_t logger_fn);
extern int pltCoreStart(const char *settings, int32_t tun_fd);
extern void pltCoreStop(int handle);
extern int64_t pltApplyConfig(int handle, const char *settings);
extern char *pltReadConfig(int handle);
extern void pltRefreshSockets(int handle);
extern void pltFixMobileRouting(int handle);
extern const char *pltCoreVersion();

#endif
