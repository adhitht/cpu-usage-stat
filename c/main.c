#include <stdio.h>
#include <string.h>
#include <unistd.h>

int *get_CPU_times(char *proc_path) {
  static int result[2] = {0};
  char file_path[256];
  snprintf(file_path, sizeof(file_path), "%s/stat", proc_path);

  FILE *file = fopen(file_path, "r");
  if (file == NULL) {
    perror("get_CPU_usage: Error opening file");
    return result;
  }

  char line[256];
  int stats[10];
  int idle_time = 0;
  int total_time = 0;

  if (fgets(line, sizeof(line), file)) { // Read first line
    sscanf(line, "cpu  %d %d %d %d %d %d %d %d %d %d", &stats[0], &stats[1],
           &stats[2], &stats[3], &stats[4], &stats[5], &stats[6], &stats[7],
           &stats[8], &stats[9]);
  }

  idle_time = stats[3];
  for (int i = 0; i < 10; i++) {
    total_time += stats[i];
  }

  fclose(file);

  result[0] = idle_time;
  result[1] = total_time;

  return result;
}

double get_CPU_usage(char *proc_path) {
  // idle_time, total_time
  int *result1 = get_CPU_times("/proc");
  int idle1 = result1[0], total1 = result1[1];
  
  usleep(1000000);
  
  int *result2 = get_CPU_times("/proc");
  int idle2 = result2[0], total2 = result2[1];

  double delta_idle_time = idle2 - idle1;
  double delta_total_time = total2 - total1;

  return (1 - delta_idle_time / delta_total_time) * 100;
}

int main() {
  char *proc_path = "/proc";
  double usage = get_CPU_usage(proc_path);

  printf("CPU Usage: %f\n", usage);
}
