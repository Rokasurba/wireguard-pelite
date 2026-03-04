// PeliteKit.swift — Pelite VPN custom WireGuard extensions

import Foundation

internal enum PeliteAdapterInfo {
    static let vendor = "Pelite"
    static let buildTag = "vpn-\(Int(Date().timeIntervalSince1970) % 100000)"
}
