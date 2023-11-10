const chokidar = require('chokidar');

interface FileWatcherOptions {
  persistent?: boolean;
}

type FileAddCallback = (filePath: string) => void;
type FileDeleteCallback = (filePath: string) => void;

export function watchFolder(folderPath: string, onFileAdd: FileAddCallback, onFileDelete: FileDeleteCallback, options: FileWatcherOptions = {}): void {
  const watcher = chokidar.watch(folderPath, {
    persistent: options.persistent ?? true,
  });

  watcher.on('add', onFileAdd);
  watcher.on('unlink', onFileDelete);
}