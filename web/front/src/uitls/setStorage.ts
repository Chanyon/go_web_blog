export function setStorage(key: string, res: unknown) {
  window.sessionStorage.setItem(key,JSON.stringify(res));
}

export function getStorage(key: string) {
  const val = window.sessionStorage.getItem(key)!;
  return JSON.parse(val);
}