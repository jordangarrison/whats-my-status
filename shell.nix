let
  unstable = import
    (fetchTarball "https://nixos.org/channels/nixos-unstable/nixexprs.tar.xz")
    { };
in { nixpkgs ? import <nixpkgs> { } }:
with nixpkgs;
mkShell {
  buildInputs = [ unstable.go unstable.go-task dbus docker gcc libGL ]
    ++ (if nixpkgs.stdenv.isLinux then [
      libcap
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
    ] else
      [ ]);
  shellHook = ''
    ${(if nixpkgs.stdenv.isDarwin then "xcode-select --install" else "")}
    go install fyne.io/fyne/v2/cmd/fyne@latest
    go install github.com/fyne-io/fyne-cross@latest
    PATH=$GOPATH/bin:/usr/bin:$PATH
  '';
}
