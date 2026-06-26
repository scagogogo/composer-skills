#!/usr/bin/env python3
"""
Generate feature diagrams for Composer Skills project.
Uses matplotlib to render beautiful tree-style feature diagrams.
"""

import matplotlib
matplotlib.use('Agg')
import matplotlib.pyplot as plt
from matplotlib.patches import FancyBboxPatch, FancyArrowPatch, Circle
import numpy as np

# ── Color palette ──
C = {
    'root':       '#1a1a2e',
    'sdk1':       '#0f3460',
    'sdk2':       '#533483',
    'cli':        '#c0392b',
    'leaf1':      '#2980b9',
    'leaf2':      '#8e44ad',
    'leaf_cli':   '#e74c3c',
    'leaf_infra': '#2c3e50',
    'bg':         '#f8f9fa',
    'white':      '#ffffff',
    'text_dark':  '#2c3e50',
    'green':      '#27ae60',
    'orange':     '#f39c12',
    'teal':       '#1abc9c',
    'dark_blue':  '#2c3e50',
    'light_blue': '#3498db',
    'purple':     '#9b59b6',
    'red':        '#e74c3c',
    'yellow':     '#f1c40f',
}


def draw_box(ax, x, y, w, h, text, color, text_color='#ffffff',
             fontsize=10, bold=False, alpha=1.0, radius=0.15,
             shadow=False, border_color=None, border_width=0.5):
    box = FancyBboxPatch(
        (x - w/2, y - h/2), w, h,
        boxstyle=f"round,pad=0.03,rounding_size={radius}",
        facecolor=color, edgecolor=border_color or color,
        linewidth=border_width, alpha=alpha, zorder=3
    )
    if shadow:
        sb = FancyBboxPatch(
            (x - w/2 + 0.04, y - h/2 - 0.04), w, h,
            boxstyle=f"round,pad=0.03,rounding_size={radius}",
            facecolor='#00000018', edgecolor='none', alpha=0.18, zorder=2
        )
        ax.add_patch(sb)
    ax.add_patch(box)
    ax.text(x, y, text, ha='center', va='center', fontsize=fontsize,
            color=text_color, weight='bold' if bold else 'normal', zorder=4,
            linespacing=1.3)


def draw_line(ax, x1, y1, x2, y2, color='#95a5a6', lw=1.5, style='-'):
    ax.plot([x1, x2], [y1, y2], color=color, linewidth=lw, alpha=0.5, zorder=1, linestyle=style)


def draw_arrow(ax, x1, y1, x2, y2, color='#95a5a6', lw=2):
    ax.annotate('', xy=(x2, y2), xytext=(x1, y1),
                arrowprops=dict(arrowstyle='->', color=color, lw=lw, alpha=0.6),
                zorder=1)


# ══════════════════════════════════════════════════════════════
# 1. Feature Mind Map — Clean tree with 3 branches
# ══════════════════════════════════════════════════════════════
def generate_feature_mindmap():
    fig, ax = plt.subplots(1, 1, figsize=(28, 20))
    ax.set_xlim(-0.5, 28.5)
    ax.set_ylim(-1, 21)
    ax.set_aspect('equal')
    ax.axis('off')
    fig.patch.set_facecolor(C['bg'])

    # Title
    ax.text(14, 20.2, 'Composer Skills  -  Feature Tree', ha='center', va='center',
            fontsize=24, color=C['root'], weight='bold')
    ax.text(14, 19.5, 'A comprehensive Go SDK and CLI for the PHP Composer ecosystem',
            ha='center', va='center', fontsize=11, color='#7f8c8d', style='italic')

    # ── Root ──
    rx, ry = 14, 18
    draw_box(ax, rx, ry, 6, 0.9, 'Composer Skills', C['root'],
             fontsize=17, bold=True, shadow=True)

    # ── Level 1: Three main branches ──
    branches = [
        (4.5, 15, 'Packagist API SDK', C['sdk1']),
        (14, 15, 'Composer CLI SDK\n(234 methods)', C['sdk2']),
        (23.5, 15, 'CLI Tool\n(50+ commands)', C['cli']),
    ]
    for bx, by, txt, col in branches:
        draw_box(ax, bx, by, 4.5, 1.0, txt, col, fontsize=12, bold=True, shadow=True)
        draw_line(ax, rx, ry - 0.45, bx, by + 0.5, color=col, lw=2.5)

    # ── Packagist API SDK leaves ──
    pkg_groups = [
        ('Package Info', 'GetPackage\nGetPackageStats\nGetPackageV2\nGetPackageDev\nGetPackageChanges', 1.5, 12.5),
        ('Search', 'SearchPackages\nSearchByTags\nSearchByType', 4.5, 12.5),
        ('Statistics', 'GetStatistics', 7.5, 12.5),
        ('Security\nAdvisories', 'GetAdvisories\nGetAdvisoriesForPkg\nGetAdvisoriesSince', 1.5, 10.5),
        ('Package\nListing', 'ListPackages\nListByVendor\nListByType\nListWithData\nListPopular', 4.5, 10.5),
        ('Package\nManagement', 'CreatePackage\nEditPackage\nUpdatePackage', 7.5, 10.5),
    ]
    for title, methods, lx, ly in pkg_groups:
        draw_box(ax, lx, ly, 2.4, 1.6, '', '#e8f4fd', alpha=0.9,
                 border_color=C['leaf1'], border_width=1.5, radius=0.15)
        ax.text(lx, ly + 0.5, title, ha='center', va='center',
                fontsize=9, color=C['leaf1'], weight='bold', zorder=5)
        ax.text(lx, ly - 0.2, methods, ha='center', va='center',
                fontsize=6.5, color='#5dade2', zorder=5, linespacing=1.2,
                family='monospace')
        draw_line(ax, 4.5, 14.5, lx, ly + 0.8, color=C['leaf1'], lw=1.2)

    # ── Composer CLI SDK: 20 categories in 4 rows ──
    cli_cats = [
        [('Core', '10'), ('Dependencies', '16'), ('Packages', '20'), ('Audit', '10'),
         ('Project', '10'), ('Config', '12')],
        [('Validation', '14'), ('Platform', '8'), ('Repository', '18'), ('Global', '14'),
         ('Auth', '10'), ('Fund', '7')],
        [('Licenses', '4'), ('Diagnosis', '8'), ('Exec', '8'), ('Satis', '8'),
         ('Version', '5'), ('Env', '12')],
        [('composer.json', '10'), ('Archive', '6'), ('Convenience', '18'),
         ('Health', '6'), ('Completion', '4')],
    ]

    for row_idx, row in enumerate(cli_cats):
        y_pos = 12.8 - row_idx * 2.0
        n = len(row)
        total_w = n * 2.4
        start_x = 14 - total_w / 2 + 1.2
        for col_idx, (name, count) in enumerate(row):
            x_pos = start_x + col_idx * 2.4
            draw_box(ax, x_pos, y_pos, 2.1, 0.85, '', '#f5eef8', alpha=0.9,
                     border_color=C['leaf2'], border_width=1, radius=0.12)
            ax.text(x_pos, y_pos + 0.15, name, ha='center', va='center',
                    fontsize=8, color=C['leaf2'], weight='bold', zorder=5)
            ax.text(x_pos, y_pos - 0.2, f'{count} methods', ha='center', va='center',
                    fontsize=7, color='#9b59b6', zorder=5)
            draw_line(ax, 14, 14.5, x_pos, y_pos + 0.42, color=C['leaf2'], lw=0.6)

    # ── CLI Tool leaves ──
    cli_groups = [
        ('Search', 'search query\nsearch by-tag\nsearch by-type', 21, 12.5),
        ('Package', 'package info\npackage stats', 23.5, 12.5),
        ('Security', 'advisories\nadvisories-for-pkg', 26, 12.5),
        ('Repository', 'repo stats\nrepo list', 21, 10.5),
        ('Local Ops', 'install / update\naudit / outdated\nwhy / fund', 23.5, 10.5),
        ('Global', 'global require\nglobal update', 26, 10.5),
    ]
    for title, cmds, lx, ly in cli_groups:
        draw_box(ax, lx, ly, 2.2, 1.5, '', '#fdedec', alpha=0.9,
                 border_color=C['leaf_cli'], border_width=1.5, radius=0.12)
        ax.text(lx, ly + 0.45, title, ha='center', va='center',
                fontsize=9, color=C['leaf_cli'], weight='bold', zorder=5)
        ax.text(lx, ly - 0.15, cmds, ha='center', va='center',
                fontsize=6.5, color='#c0392b', zorder=5, linespacing=1.2,
                family='monospace')
        draw_line(ax, 23.5, 14.5, lx, ly + 0.75, color=C['leaf_cli'], lw=1.2)

    # ── Foundation bar ──
    draw_box(ax, 14, 3.5, 27, 1.8, '', '#ecf0f1', alpha=0.7,
             border_color='#bdc3c7', border_width=1.5, radius=0.3)
    ax.text(14, 4.1, 'Foundation Layer', ha='center', va='center',
            fontsize=15, color=C['root'], weight='bold')
    founds = [
        ('Domain Models\n(Package, Advisory,\nStatistics, Version...)', 4, 3.1),
        ('Detector\n(Cross-OS\nDetection)', 8.5, 3.1),
        ('Installer\n(Auto-Install\nPHP + Composer)', 13, 3.1),
        ('Utilities\n(FS, HTTP,\nMock Helpers)', 17.5, 3.1),
        ('Composerutils\n(Shared Utils)', 22, 3.1),
    ]
    for txt, fx, fy in founds:
        draw_box(ax, fx, fy, 3.5, 0.7, txt, C['leaf_infra'], fontsize=7.5,
                 text_color='white', radius=0.1)

    # ── Stats box (bottom-left) ──
    draw_box(ax, 3, 6, 5.5, 2.5, '', '#ffffff', alpha=0.9,
             border_color=C['sdk1'], border_width=2, radius=0.15, shadow=True)
    ax.text(3, 7.0, 'Project Stats', ha='center', va='center',
            fontsize=12, color=C['root'], weight='bold')
    stats = [
        ('20 Packagist API Methods', C['sdk1']),
        ('234 Composer CLI Methods', C['sdk2']),
        ('20 Method Categories', C['leaf2']),
        ('50+ CLI Subcommands', C['cli']),
        ('450+ Test Cases', C['green']),
        ('12 Documentation Guides', '#2980b9'),
    ]
    for i, (txt, col) in enumerate(stats):
        ax.text(3, 6.3 - i * 0.35, txt, ha='center', va='center',
                fontsize=9, color=col, weight='bold')

    plt.tight_layout(pad=0.3)
    plt.savefig('/home/cc11001100/github/scagogogo/composer-skills/docs/images/feature-mindmap.png',
                dpi=150, bbox_inches='tight', facecolor=C['bg'])
    print("Feature mind map saved!")


# ══════════════════════════════════════════════════════════════
# 2. Architecture Diagram — Enhanced with data flow arrows
# ══════════════════════════════════════════════════════════════
def generate_architecture_diagram():
    fig, ax = plt.subplots(1, 1, figsize=(18, 11))
    ax.set_xlim(0, 18)
    ax.set_ylim(0, 11)
    ax.set_aspect('equal')
    ax.axis('off')
    fig.patch.set_facecolor('#ffffff')

    ax.text(9, 10.5, 'Three-Layer Architecture', ha='center', va='center',
            fontsize=22, color=C['root'], weight='bold')
    ax.text(9, 10.0, 'Data flows from top to bottom, capabilities flow from bottom to top',
            ha='center', va='center', fontsize=10, color='#7f8c8d', style='italic')

    # Layer 1: Docs
    draw_box(ax, 9, 8.8, 15, 1.0, '', '#eaf2f8', alpha=0.9,
             border_color='#2980b9', border_width=2, radius=0.2)
    ax.text(9, 9.0, 'Skills Documentation Layer', ha='center', va='center',
            fontsize=14, color='#2980b9', weight='bold')
    ax.text(9, 8.6, 'Progressive disclosure guides (12 guides) | API reference | Examples',
            ha='center', va='center', fontsize=10, color='#5dade2')

    # Layer 2: CLI
    draw_box(ax, 9, 7.2, 15, 1.0, '', '#fdedec', alpha=0.9,
             border_color=C['cli'], border_width=2, radius=0.2)
    ax.text(9, 7.4, 'CLI Tool Layer (Cobra)', ha='center', va='center',
            fontsize=14, color=C['cli'], weight='bold')
    ax.text(9, 7.0, '50+ subcommands | composer-skills command-line interface',
            ha='center', va='center', fontsize=10, color='#e74c3c')

    # Layer 3: Two SDKs
    draw_box(ax, 4.5, 5.4, 7, 1.2, '', '#eaf2f8', alpha=0.9,
             border_color=C['sdk1'], border_width=2, radius=0.2)
    ax.text(4.5, 5.7, 'Packagist API SDK', ha='center', va='center',
            fontsize=13, color=C['sdk1'], weight='bold')
    ax.text(4.5, 5.3, 'pkg/client | pkg/repository', ha='center', va='center',
            fontsize=9, color='#5dade2')
    ax.text(4.5, 5.0, '20 methods | Pure Go (No PHP)', ha='center', va='center',
            fontsize=9, color='#7f8c8d')

    draw_box(ax, 13.5, 5.4, 7, 1.2, '', '#f5eef8', alpha=0.9,
             border_color=C['sdk2'], border_width=2, radius=0.2)
    ax.text(13.5, 5.7, 'Composer CLI SDK', ha='center', va='center',
            fontsize=13, color=C['sdk2'], weight='bold')
    ax.text(13.5, 5.3, 'pkg/composer', ha='center', va='center',
            fontsize=9, color='#9b59b6')
    ax.text(13.5, 5.0, '234 methods | 20 categories', ha='center', va='center',
            fontsize=9, color='#7f8c8d')

    # Layer 4: Foundation
    draw_box(ax, 9, 3.1, 15, 1.0, '', '#f0f0f0', alpha=0.9,
             border_color='#95a5a6', border_width=2, radius=0.2)
    ax.text(9, 3.3, 'Foundation Layer', ha='center', va='center',
            fontsize=14, color=C['leaf_infra'], weight='bold')
    ax.text(9, 2.9, 'pkg/domain | pkg/detector | pkg/installer | pkg/composerutils',
            ha='center', va='center', fontsize=10, color='#95a5a6')

    # Arrows between layers
    arrow_kw = dict(arrowstyle='->', color='#bdc3c7', lw=2)
    for xy, xytext in [
        ((9, 8.3), (9, 7.7)),
        ((4.5, 6.7), (4.5, 6.0)),
        ((13.5, 6.7), (13.5, 6.0)),
        ((4.5, 4.8), (4.5, 3.6)),
        ((13.5, 4.8), (13.5, 3.6)),
    ]:
        ax.annotate('', xy=xy, xytext=xytext, arrowprops=arrow_kw)

    # External connections
    # Packagist cloud
    draw_box(ax, 4.5, 1.5, 3.5, 0.8, 'Packagist.org', '#3498db',
             fontsize=11, bold=True, shadow=True)
    ax.text(4.5, 1.0, 'REST API (HTTPS)', ha='center', va='center',
            fontsize=8, color='#7f8c8d')
    draw_arrow(ax, 4.5, 2.6, 4.5, 1.9, color=C['sdk1'], lw=2)

    # PHP/Composer binary
    draw_box(ax, 13.5, 1.5, 3.5, 0.8, 'PHP + Composer', '#9b59b6',
             fontsize=11, bold=True, shadow=True)
    ax.text(13.5, 1.0, 'Local Binary (CLI)', ha='center', va='center',
            fontsize=8, color='#7f8c8d')
    draw_arrow(ax, 13.5, 2.6, 13.5, 1.9, color=C['sdk2'], lw=2)

    plt.tight_layout(pad=0.5)
    plt.savefig('/home/cc11001100/github/scagogogo/composer-skills/docs/images/architecture.png',
                dpi=150, bbox_inches='tight', facecolor='#ffffff')
    print("Architecture diagram saved!")


# ══════════════════════════════════════════════════════════════
# 3. SDK Comparison Banner
# ══════════════════════════════════════════════════════════════
def generate_sdk_comparison():
    fig, ax = plt.subplots(1, 1, figsize=(16, 6))
    ax.set_xlim(0, 16)
    ax.set_ylim(0, 6)
    ax.set_aspect('equal')
    ax.axis('off')
    fig.patch.set_facecolor('#ffffff')

    # Left: Packagist
    draw_box(ax, 4, 3, 6.5, 4.5, '', '#f0f7ff', alpha=1.0,
             border_color=C['sdk1'], border_width=2.5, radius=0.3, shadow=True)
    ax.text(4, 4.9, 'Packagist API SDK', ha='center', va='center',
            fontsize=15, color=C['sdk1'], weight='bold')
    ax.text(4, 4.35, 'pkg/client  |  pkg/repository', ha='center', va='center',
            fontsize=10, color='#5dade2')

    left_feats = [
        'Search packages & browse registry',
        'Get package details & statistics',
        'Security advisories & CVE tracking',
        'List packages by vendor / type',
        'Pure Go -- no PHP required',
    ]
    for i, f in enumerate(left_feats):
        ax.text(4, 3.6 - i * 0.45, '  ' + f, ha='center', va='center',
                fontsize=9.5, color=C['text_dark'])

    # Right: Composer CLI
    draw_box(ax, 12, 3, 6.5, 4.5, '', '#faf0ff', alpha=1.0,
             border_color=C['sdk2'], border_width=2.5, radius=0.3, shadow=True)
    ax.text(12, 4.9, 'Composer CLI SDK', ha='center', va='center',
            fontsize=15, color=C['sdk2'], weight='bold')
    ax.text(12, 4.35, 'pkg/composer', ha='center', va='center',
            fontsize=10, color='#9b59b6')

    right_feats = [
        '234 methods across 20 categories',
        'Install / update / require / remove',
        'Security audit & vulnerability scan',
        'Project management & scripts',
        'Requires PHP 7.4+ & Composer 2.0+',
    ]
    for i, f in enumerate(right_feats):
        ax.text(12, 3.6 - i * 0.45, '  ' + f, ha='center', va='center',
                fontsize=9.5, color=C['text_dark'])

    # VS circle
    circ = plt.Circle((8, 3), 0.55, color=C['root'], alpha=0.9, zorder=5)
    ax.add_patch(circ)
    ax.text(8, 3, 'VS', ha='center', va='center', fontsize=12,
            color='white', weight='bold', zorder=6)

    plt.tight_layout(pad=0.5)
    plt.savefig('/home/cc11001100/github/scagogogo/composer-skills/docs/images/sdk-comparison.png',
                dpi=150, bbox_inches='tight', facecolor='#ffffff')
    print("SDK comparison saved!")


# ══════════════════════════════════════════════════════════════
# 4. Header Badge
# ══════════════════════════════════════════════════════════════
def generate_header_badge():
    fig, ax = plt.subplots(1, 1, figsize=(14, 3))
    ax.set_xlim(0, 14)
    ax.set_ylim(0, 3)
    ax.set_aspect('equal')
    ax.axis('off')

    for i in range(140):
        x = i / 10
        r = 0.06 + 0.04 * (i / 140)
        b = 0.25 + 0.15 * (i / 140)
        ax.axvspan(x, x + 0.1, facecolor=(r, 0.15, b), alpha=1.0)

    ax.text(7, 2.0, 'Composer Skills', ha='center', va='center',
            fontsize=28, color='white', weight='bold')
    ax.text(7, 1.1, 'Go SDK & CLI for the PHP Composer Ecosystem',
            ha='center', va='center', fontsize=13, color='#a8d8ea', style='italic')

    badges = [
        ('234', 'SDK Methods', 2.5),
        ('20', 'API Methods', 5.5),
        ('50+', 'CLI Commands', 8.5),
        ('450+', 'Tests', 11.5),
    ]
    for num, label, x in badges:
        draw_box(ax, x, 0.4, 2.2, 0.55, '', '#ffffff15',
                 border_color='#ffffff40', border_width=1, radius=0.1)
        ax.text(x - 0.3, 0.4, num, ha='center', va='center',
                fontsize=11, color=C['orange'], weight='bold')
        ax.text(x + 0.5, 0.4, label, ha='center', va='center',
                fontsize=8, color='#bdc3c7')

    plt.savefig('/home/cc11001100/github/scagogogo/composer-skills/docs/images/header-badge.png',
                dpi=150, bbox_inches='tight', facecolor='#1a1a2e')
    print("Header badge saved!")


# ══════════════════════════════════════════════════════════════
# 5. Auto-Install Flow Diagram
# ══════════════════════════════════════════════════════════════
def generate_auto_install_flow():
    fig, ax = plt.subplots(1, 1, figsize=(16, 10))
    ax.set_xlim(0, 16)
    ax.set_ylim(0, 10)
    ax.set_aspect('equal')
    ax.axis('off')
    fig.patch.set_facecolor('#ffffff')

    ax.text(8, 9.5, 'Auto-Install Flow', ha='center', va='center',
            fontsize=22, color=C['root'], weight='bold')
    ax.text(8, 9.0, 'Zero-config Composer setup — detect, install, verify, ready',
            ha='center', va='center', fontsize=11, color='#7f8c8d', style='italic')

    # Flow steps
    steps = [
        (2, 7.5, '1. Detect', 'Is Composer\ninstalled?', C['sdk1'], '#eaf2f8'),
        (5.5, 7.5, '2. Check PHP', 'Is PHP\navailable?', C['sdk2'], '#f5eef8'),
        (9, 7.5, '3. Install PHP', 'Auto-install\nvia pkg manager', C['orange'], '#fef9e7'),
        (12.5, 7.5, '4. Install\nComposer', 'pkg manager\nor direct download', C['cli'], '#fdedec'),
        (8, 4.5, '5. Verify', 'Detect &\nvalidate version', C['green'], '#eafaf1'),
        (8, 2.0, '6. Ready!', 'Composer\ninstance created', C['root'], '#f8f9fa'),
    ]

    for x, y, title, desc, title_color, bg_color in steps:
        draw_box(ax, x, y, 2.8, 1.6, '', bg_color, alpha=0.95,
                 border_color=title_color, border_width=2, radius=0.2, shadow=True)
        ax.text(x, y + 0.4, title, ha='center', va='center',
                fontsize=11, color=title_color, weight='bold', zorder=5)
        ax.text(x, y - 0.3, desc, ha='center', va='center',
                fontsize=9, color='#555555', zorder=5, linespacing=1.3)

    # Arrows
    draw_arrow(ax, 3.4, 7.5, 4.1, 7.5, color='#bdc3c7', lw=2)
    draw_arrow(ax, 6.9, 7.5, 7.6, 7.5, color='#bdc3c7', lw=2)
    draw_arrow(ax, 10.4, 7.5, 11.1, 7.5, color='#bdc3c7', lw=2)
    draw_arrow(ax, 12.5, 6.7, 8, 5.3, color='#bdc3c7', lw=2)
    draw_arrow(ax, 8, 3.7, 8, 2.8, color='#bdc3c7', lw=2)

    # "Already installed" shortcut
    draw_box(ax, 2, 4.5, 2.8, 1.0, 'Already\ninstalled!', C['green'],
             fontsize=9, bold=True, text_color='white', shadow=True)
    draw_arrow(ax, 2, 6.7, 2, 5.0, color=C['green'], lw=2)
    draw_line(ax, 3.4, 4.5, 6.6, 2.0, color=C['green'], lw=2, style='--')

    # "PHP available" shortcut
    draw_box(ax, 5.5, 5.5, 2.2, 0.7, 'PHP OK!', C['teal'],
             fontsize=9, bold=True, text_color='white')
    draw_arrow(ax, 5.5, 6.7, 5.5, 5.85, color=C['teal'], lw=1.5)
    draw_line(ax, 6.6, 5.5, 11.1, 7.5, color=C['teal'], lw=1.5, style='--')

    plt.tight_layout(pad=0.5)
    plt.savefig('/home/cc11001100/github/scagogogo/composer-skills/docs/images/auto-install-flow.png',
                dpi=150, bbox_inches='tight', facecolor='#ffffff')
    print("Auto-install flow diagram saved!")


# ══════════════════════════════════════════════════════════════
# 6. Platform Support Matrix
# ══════════════════════════════════════════════════════════════
def generate_platform_matrix():
    fig, ax = plt.subplots(1, 1, figsize=(14, 7))
    ax.set_xlim(0, 14)
    ax.set_ylim(0, 7)
    ax.set_aspect('equal')
    ax.axis('off')
    fig.patch.set_facecolor('#ffffff')

    ax.text(7, 6.5, 'Cross-Platform Support', ha='center', va='center',
            fontsize=20, color=C['root'], weight='bold')

    # Column headers
    platforms = [
        (3.5, 'Linux', '#e74c3c', ['Ubuntu / Debian', 'CentOS / RHEL', 'Fedora', 'Arch / Manjaro', 'Alpine', 'openSUSE', 'Gentoo']),
        (7, 'macOS', '#2980b9', ['Homebrew', 'Direct Download']),
        (10.5, 'Windows', '#27ae60', ['Direct Download', 'PATH Setup']),
    ]

    for px, name, color, distros in platforms:
        # Platform header
        draw_box(ax, px, 5.5, 3, 0.8, name, color,
                 fontsize=14, bold=True, text_color='white', shadow=True)

        # Features
        features = [
            'Composer Detection',
            'Auto-Install',
            'PHP Auto-Install',
            'Package Manager',
        ]
        for i, feat in enumerate(features):
            fy = 4.5 - i * 0.9
            draw_box(ax, px, fy, 2.8, 0.6, '', '#f8f9fa', alpha=0.9,
                     border_color='#ddd', border_width=1, radius=0.1)
            ax.text(px, fy, feat, ha='center', va='center',
                    fontsize=9, color=C['text_dark'], zorder=5)

            # Checkmark
            if name == 'Linux':
                if feat == 'Package Manager':
                    checks = 'apt/dnf/pacman/apk/zypper'
                else:
                    checks = '✓'
            elif name == 'macOS':
                if feat == 'Package Manager':
                    checks = 'Homebrew'
                else:
                    checks = '✓'
            else:
                if feat == 'PHP Auto-Install':
                    checks = '—'
                elif feat == 'Package Manager':
                    checks = '—'
                else:
                    checks = '✓'

            check_color = C['green'] if checks == '✓' else ('#e67e22' if checks == '—' else '#7f8c8d')
            ax.text(px + 1.1, fy, checks, ha='center', va='center',
                    fontsize=8, color=check_color, weight='bold', zorder=5)

    # Legend
    ax.text(7, 0.5, '✓ = Supported    — = Not applicable    Package Manager = Primary installation method',
            ha='center', va='center', fontsize=9, color='#7f8c8d', style='italic')

    plt.tight_layout(pad=0.5)
    plt.savefig('/home/cc11001100/github/scagogogo/composer-skills/docs/images/platform-matrix.png',
                dpi=150, bbox_inches='tight', facecolor='#ffffff')
    print("Platform matrix saved!")


# ══════════════════════════════════════════════════════════════
# 7. Security Feature Diagram
# ══════════════════════════════════════════════════════════════
def generate_security_diagram():
    fig, ax = plt.subplots(1, 1, figsize=(14, 8))
    ax.set_xlim(0, 14)
    ax.set_ylim(0, 8)
    ax.set_aspect('equal')
    ax.axis('off')
    fig.patch.set_facecolor('#ffffff')

    ax.text(7, 7.5, 'Security-First Design', ha='center', va='center',
            fontsize=20, color=C['root'], weight='bold')

    # Center shield
    draw_box(ax, 7, 5.5, 3.5, 1.2, 'Security\nAudit', '#c0392b',
             fontsize=14, bold=True, text_color='white', shadow=True)

    # Surrounding features
    features = [
        (2.5, 6.0, 'Vulnerability\nScan', '#e74c3c', 'HasVulnerabilities\nGetHighSeverityVulns'),
        (11.5, 6.0, 'Packagist\nAdvisories', '#e67e22', 'GetSecurityAdvisories\nGetAdvisoriesForPkg'),
        (2.5, 3.5, 'Schema\nValidation', '#2980b9', 'Validate\nValidateStrict\nValidateSchema'),
        (11.5, 3.5, 'Platform\nRequirements', '#27ae60', 'CheckPlatform\nGetPHPVersion\nHasExtension'),
        (7, 2.5, 'License\nCompliance', '#8e44ad', 'Licenses\nCheckLicenses\nLicensesWithFormat'),
    ]

    for fx, fy, title, color, methods in features:
        draw_box(ax, fx, fy, 3, 1.4, '', '#f8f9fa', alpha=0.95,
                 border_color=color, border_width=2, radius=0.2, shadow=True)
        ax.text(fx, fy + 0.35, title, ha='center', va='center',
                fontsize=10, color=color, weight='bold', zorder=5)
        ax.text(fx, fy - 0.3, methods, ha='center', va='center',
                fontsize=7, color='#555555', zorder=5, linespacing=1.2, family='monospace')
        draw_line(ax, 7, 4.9, fx, fy + 0.7, color=color, lw=1.5)

    plt.tight_layout(pad=0.5)
    plt.savefig('/home/cc11001100/github/scagogogo/composer-skills/docs/images/security-features.png',
                dpi=150, bbox_inches='tight', facecolor='#ffffff')
    print("Security features diagram saved!")


if __name__ == '__main__':
    generate_feature_mindmap()
    generate_architecture_diagram()
    generate_sdk_comparison()
    generate_header_badge()
    generate_auto_install_flow()
    generate_platform_matrix()
    generate_security_diagram()
    print("\nAll images generated successfully!")
