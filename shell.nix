{ pkgs ? import <nixpkgs> {} }:
pkgs.mkShell {
  buildInputs = with pkgs; [
    gnumake
    go
    go-tools
    libpcap
    gosec
  ];
}
