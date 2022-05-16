let
  unstable = import
    (fetchTarball "https://nixos.org/channels/nixos-unstable/nixexprs.tar.xz")
    { };
in { nixpkgs ? import <nixpkgs> { } }:
with nixpkgs;
mkShell {
  buildInputs = [
    unstable.go
    unstable.go-task
    dbus
    libGL
    libcap
    gcc
    xorg.libX11
    xorg.libX11.dev
    xorg.libX11.dev.out
    xorg.libXcursor
    xorg.libXrandr
    xorg.libXinerama
    xorg.libXi
    xorg.xinput
    xorg.libXext
    xorg.libXxf86vm
    pkg-config
  ];
  # nativeBuildInputs = [
  # ];
  shellHook = ''
    go get -u github.com/fyne-io/fyne-cross
  '';
}
