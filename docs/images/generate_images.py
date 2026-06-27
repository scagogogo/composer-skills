#!/usr/bin/env python3
"""
Generate feature diagrams for Composer Skills project.
Modern design — clean typography, subtle shadows, card-based layout.
No emoji (matplotlib can't render them). Uses text labels instead.
"""

import matplotlib
matplotlib.use('Agg')
import matplotlib.pyplot as plt
from matplotlib.patches import FancyBboxPatch, Circle
from matplotlib.colors import LinearSegmentedColormap
import numpy as np

# ── Modern Design System ──
PAL = {
    'primary':     '#4F46E5',
    'primary_lt':  '#818CF8',
    'primary_bg':  '#EEF2FF',

    'secondary':   '#7C3AED',
    'secondary_lt': '#A78BFA',
    'secondary_bg': '#F5F3FF',

    'accent':      '#E11D48',
    'accent_lt':   '#FB7185',
    'accent_bg':   '#FFF1F2',

    'success':    '#059669',
    'success_lt':'#34D399',
    'success_bg': '#ECFDF5',

    'warning':    '#D97706',
    'warning_lt': '#FBBF24',
    'warning_bg': '#FFFBEB',

    'info':       '#0284C7',
    'info_lt':   '#38BDF8',
    'info_bg':    '#F0F9FF',

    'bg':         '#FAFAFA',
    'surface':    '#FFFFFF',
    'surface2':   '#F8FAFC',
    'border':     '#E2E8F0',
    'border2':    '#CBD5E1',
    'text':       '#1E293B',
    'text2':      '#475569',
    'text3':      '#94A3B8',
}


def draw_card(ax, x, y, w, h, title='', subtitle='', color=PAL['primary'],
              bg=PAL['primary_bg'], title_fontsize=12, sub_fontsize=9,
              shadow=True, border_width=1.5, radius=0.18, bold=True,
              title_color=None, sub_color=None, zorder=3):
    """Modern card: rounded corners, accent bar, optional shadow."""
    title_color = title_color or color
    sub_color = sub_color or PAL['text2']

    if shadow:
        sb = FancyBboxPatch(
            (x - w/2 + 0.05, y - h/2 - 0.05), w, h,
            boxstyle=f"round,pad=0.02,rounding_size={radius}",
            facecolor='#00000007', edgecolor='none',
            alpha=0.10, zorder=zorder-1
        )
        ax.add_patch(sb)

    card = FancyBboxPatch(
        (x - w/2, y - h/2), w, h,
        boxstyle=f"round,pad=0.02,rounding_size={radius}",
        facecolor=bg, edgecolor=color,
        linewidth=border_width, alpha=0.95, zorder=zorder
    )
    ax.add_patch(card)

    # Accent bar at top
    bar = FancyBboxPatch(
        (x - w/2, y + h/2 - 0.06), w, 0.06,
        boxstyle=f"round,pad=0.0,rounding_size={radius}",
        facecolor=color, edgecolor='none', alpha=0.85, zorder=zorder+1
    )
    ax.add_patch(bar)

    # Text
    if subtitle:
        ax.text(x, y + 0.12, title, ha='center', va='center',
                fontsize=title_fontsize, color=title_color,
                weight='bold' if bold else 'normal', zorder=zorder+2)
        ax.text(x, y - 0.22, subtitle, ha='center', va='center',
                fontsize=sub_fontsize, color=sub_color, zorder=zorder+2,
                linespacing=1.3)
    else:
        ax.text(x, y, title, ha='center', va='center',
                fontsize=title_fontsize, color=title_color,
                weight='bold' if bold else 'normal', zorder=zorder+2)


def draw_connector(ax, x1, y1, x2, y2, color=PAL['border2'], lw=1.5,
                   arrow=True, curved=False, style='-'):
    """Clean connector line with optional arrow."""
    if arrow:
        ax.annotate('', xy=(x2, y2), xytext=(x1, y1),
                    arrowprops=dict(arrowstyle='->', color=color, lw=lw,
                                   alpha=0.5,
                                   connectionstyle='arc3,rad=0.08' if curved else 'arc3,rad=0'),
                    zorder=1)
    else:
        ax.plot([x1, x2], [y1, y2], color=color, linewidth=lw, alpha=0.35,
                linestyle=style, zorder=1)


def draw_gradient_bg(ax, xlim, ylim, c1, c2):
    """Subtle gradient background."""
    grad = np.linspace(0, 1, 256).reshape(1, -1)
    grad = np.vstack([grad]*10)
    cmap = LinearSegmentedColormap.from_list('bg', [c1, c2])
    ax.imshow(grad, aspect='auto', cmap=cmap,
              extent=[xlim[0], xlim[1], ylim[0], ylim[1]], alpha=0.04, zorder=0)


# ══════════════════════════════════════════════════════════════
# 1. Header Badge
# ══════════════════════════════════════════════════════════════
def generate_header_badge():
    fig, ax = plt.subplots(1, 1, figsize=(14, 3.5))
    ax.set_xlim(0, 14)
    ax.set_ylim(0, 3.5)
    ax.set_aspect('equal')
    ax.axis('off')

    # Deep gradient
    for i in range(280):
        x = i / 20
        t = i / 280
        r = 0.06 + 0.14 * t
        g = 0.05 + 0.06 * t
        b = 0.28 + 0.18 * t
        ax.axvspan(x, x + 0.05, facecolor=(r, g, b), alpha=1.0, zorder=0)

    # Decorative soft circles
    for cx, cy, cr in [(2.5, 2.8, 1.5), (10, 1.2, 1.3), (6, 3.2, 0.9)]:
        c = plt.Circle((cx, cy), cr, color='white', alpha=0.02, zorder=1)
        ax.add_patch(c)

    ax.text(7, 2.2, 'Composer Skills', ha='center', va='center',
            fontsize=30, color='white', weight='bold', zorder=5)
    ax.text(7, 1.5, 'Go SDK & CLI for the PHP Composer Ecosystem',
            ha='center', va='center', fontsize=12, color='#C7D2FE',
            zorder=5)

    # Stat badges
    stats = [('234', 'SDK Methods', 2.5), ('20', 'API Methods', 5.5),
             ('50+', 'CLI Commands', 8.5), ('450+', 'Tests', 11.5)]
    for num, label, x in stats:
        card = FancyBboxPatch(
            (x - 1.1, 0.25), 2.2, 0.55,
            boxstyle="round,pad=0.02,rounding_size=0.15",
            facecolor='#ffffff12', edgecolor='#ffffff25',
            linewidth=1, zorder=4
        )
        ax.add_patch(card)
        ax.text(x - 0.35, 0.52, num, ha='center', va='center',
                fontsize=13, color='#FBBF24', weight='bold', zorder=5)
        ax.text(x + 0.5, 0.52, label, ha='center', va='center',
                fontsize=8, color='#94A3B8', zorder=5)

    plt.savefig('/home/cc11001100/github/scagogogo/composer-skills/docs/images/header-badge.png',
                dpi=180, bbox_inches='tight', facecolor='#1E1B4B')
    print("Header badge saved!")


# ══════════════════════════════════════════════════════════════
# 2. Architecture
# ══════════════════════════════════════════════════════════════
def generate_architecture_diagram():
    fig, ax = plt.subplots(1, 1, figsize=(16, 10))
    ax.set_xlim(0, 16)
    ax.set_ylim(0, 10)
    ax.set_aspect('equal')
    ax.axis('off')
    fig.patch.set_facecolor(PAL['bg'])
    draw_gradient_bg(ax, (0, 16), (0, 10), '#EEF2FF', '#F5F3FF')

    ax.text(8, 9.5, 'Architecture', ha='center', va='center',
            fontsize=24, color=PAL['text'], weight='bold')
    ax.text(8, 9.05, 'Docs > CLI > SDKs > Foundation',
            ha='center', va='center', fontsize=10, color=PAL['text3'])

    draw_card(ax, 8, 7.8, 12, 0.9, 'Documentation Layer',
              'Progressive guides (12) | API reference | 15+ examples',
              color=PAL['info'], bg=PAL['info_bg'], title_fontsize=13)

    draw_card(ax, 8, 6.2, 12, 0.9, 'CLI Tool (Cobra)',
              '50+ subcommands | composer-skills CLI',
              color=PAL['accent'], bg=PAL['accent_bg'], title_fontsize=13)

    draw_card(ax, 4, 4.2, 5.5, 1.3, 'Packagist API SDK',
              'pkg/client | pkg/repository\n20 methods | Pure Go (No PHP)',
              color=PAL['primary'], bg=PAL['primary_bg'], title_fontsize=12)

    draw_card(ax, 12, 4.2, 5.5, 1.3, 'Composer CLI SDK',
              'pkg/composer\n234 methods | 20 categories',
              color=PAL['secondary'], bg=PAL['secondary_bg'], title_fontsize=12)

    draw_card(ax, 8, 2.0, 12, 0.9, 'Foundation Layer',
              'Domain | Detector | Installer | Utilities',
              color=PAL['text2'], bg=PAL['surface2'], title_fontsize=13)

    # Flow arrows
    draw_connector(ax, 8, 7.35, 8, 6.65, color=PAL['info'], lw=2.5)
    draw_connector(ax, 5.5, 5.7, 4.5, 4.85, color=PAL['primary'], lw=2)
    draw_connector(ax, 10.5, 5.7, 11.5, 4.85, color=PAL['secondary'], lw=2)
    draw_connector(ax, 4, 3.55, 5.5, 2.45, color=PAL['border2'], lw=1.5, curved=True)
    draw_connector(ax, 12, 3.55, 10.5, 2.45, color=PAL['border2'], lw=1.5, curved=True)

    # External — Packagist.org
    draw_card(ax, 4, 0.6, 3.3, 0.65, 'Packagist.org',
              'REST API (HTTPS)',
              color=PAL['info_lt'], bg=PAL['info_bg'], title_fontsize=10, shadow=False)
    draw_connector(ax, 4, 1.5, 4, 3.55, color=PAL['info'], lw=1.5, style='--')

    # External — PHP
    draw_card(ax, 12, 0.6, 3.3, 0.65, 'PHP + Composer',
              'Local binary (CLI)',
              color=PAL['secondary_lt'], bg=PAL['secondary_bg'], title_fontsize=10, shadow=False)
    draw_connector(ax, 12, 1.5, 12, 3.55, color=PAL['secondary'], lw=1.5, style='--')

    plt.tight_layout(pad=0.5)
    plt.savefig('/home/cc11001100/github/scagogogo/composer-skills/docs/images/architecture.png',
                dpi=150, bbox_inches='tight', facecolor=PAL['bg'])
    print("Architecture saved!")


# ══════════════════════════════════════════════════════════════
# 3. Feature Mind Map
# ══════════════════════════════════════════════════════════════
def generate_feature_mindmap():
    fig, ax = plt.subplots(1, 1, figsize=(26, 22))
    ax.set_xlim(-0.5, 26.5)
    ax.set_ylim(-0.5, 22.5)
    ax.set_aspect('equal')
    ax.axis('off')
    fig.patch.set_facecolor(PAL['bg'])
    draw_gradient_bg(ax, (-0.5, 26.5), (-0.5, 22.5), '#EEF2FF', '#F5F3FF')

    ax.text(13, 22, 'Feature Map', ha='center', va='center',
            fontsize=26, color=PAL['text'], weight='bold')
    ax.text(13, 21.4, 'Two SDKs + CLI -- everything for the PHP/Composer ecosystem',
            ha='center', va='center', fontsize=11, color=PAL['text3'])

    # Root
    draw_card(ax, 13, 19.5, 5, 0.85, 'Composer Skills',
              color=PAL['primary'], bg=PAL['primary_bg'], title_fontsize=16, shadow=True)

    # 3 branches
    draw_card(ax, 4.5, 17, 4, 0.8, 'Packagist API', '20 methods | Pure Go',
              color=PAL['info'], bg=PAL['info_bg'], title_fontsize=11, shadow=True)
    draw_card(ax, 13, 17, 4, 0.8, 'Composer CLI', '234 methods | 20 categories',
              color=PAL['secondary'], bg=PAL['secondary_bg'], title_fontsize=11, shadow=True)
    draw_card(ax, 21.5, 17, 4, 0.8, 'CLI Tool', '50+ subcommands',
              color=PAL['accent'], bg=PAL['accent_bg'], title_fontsize=11, shadow=True)

    draw_connector(ax, 13, 19.07, 4.5, 17.4, color=PAL['info_lt'], lw=2.5)
    draw_connector(ax, 13, 19.07, 13, 17.4, color=PAL['secondary_lt'], lw=2.5)
    draw_connector(ax, 13, 19.07, 21.5, 17.4, color=PAL['accent_lt'], lw=2.5)

    # Packagist sub-cards
    api_feats = [
        ('Package Info', 'GetPackage/Stats/V2\nDev/Changes', PAL['info'], PAL['info_bg']),
        ('Search', 'SearchPackages\nByTags/ByType', PAL['info'], PAL['info_bg']),
        ('Statistics', 'GetStatistics', PAL['info'], PAL['info_bg']),
        ('Security', 'Advisories/ForPkg\nSince', PAL['warning'], PAL['warning_bg']),
        ('Listing', 'List/ByVendor\nByType/Popular', PAL['info'], PAL['info_bg']),
        ('Management', 'Create/Edit\nUpdate', PAL['primary'], PAL['primary_bg']),
    ]
    for i, (t, s, c, b) in enumerate(api_feats):
        row = i // 3; col = i % 3
        lx = 1.5 + col * 3; ly = 14.7 - row * 2.2
        draw_card(ax, lx, ly, 2.6, 1.4, t, s,
                  color=c, bg=b, title_fontsize=9, sub_fontsize=6.5,
                  shadow=False, border_width=1.2)
        draw_connector(ax, 4.5, 16.6, lx, ly + 0.7, color=PAL['info_lt'], lw=1, arrow=False)

    # Composer CLI sub-cards
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
    for ri, row in enumerate(cli_cats):
        yp = 14.7 - ri * 2.1
        n = len(row)
        tw = n * 2.3; sx = 13 - tw/2 + 1.15
        for ci, (nm, ct) in enumerate(row):
            xp = sx + ci * 2.3
            draw_card(ax, xp, yp, 2.0, 0.8, nm, f'{ct} methods',
                      color=PAL['secondary'], bg=PAL['secondary_bg'],
                      title_fontsize=8, sub_fontsize=6.5, shadow=False, border_width=1)
            draw_connector(ax, 13, 16.6, xp, yp + 0.4, color=PAL['secondary_lt'], lw=0.5, arrow=False)

    # CLI Tool sub-cards
    cli_feats = [
        ('Search', 'query/by-tag\n/by-type', PAL['accent'], PAL['accent_bg']),
        ('Package', 'info/stats', PAL['accent'], PAL['accent_bg']),
        ('Security', 'advisories\n/for-pkg', PAL['warning'], PAL['warning_bg']),
        ('Repository', 'stats/list', PAL['accent'], PAL['accent_bg']),
        ('Local', 'install/update\naudit/outdated', PAL['secondary'], PAL['secondary_bg']),
        ('Global', 'require\nupdate', PAL['primary'], PAL['primary_bg']),
    ]
    for i, (t, s, c, b) in enumerate(cli_feats):
        row = i // 3; col = i % 3
        lx = 20 + col * 3; ly = 14.7 - row * 2.2
        draw_card(ax, lx, ly, 2.6, 1.4, t, s,
                  color=c, bg=b, title_fontsize=9, sub_fontsize=6.5,
                  shadow=False, border_width=1.2)
        draw_connector(ax, 21.5, 16.6, lx, ly + 0.7, color=PAL['accent_lt'], lw=1, arrow=False)

    # Foundation bar
    fb = FancyBboxPatch(
        (0.2, 2.6), 26, 1.8,
        boxstyle="round,pad=0.04,rounding_size=0.35",
        facecolor=PAL['surface'], edgecolor=PAL['border'],
        linewidth=2, alpha=0.7, zorder=2
    )
    ax.add_patch(fb)
    ax.text(13, 4.0, 'Foundation Layer', ha='center', va='center',
            fontsize=16, color=PAL['text'], weight='bold', zorder=3)

    f_items = [
        ('Domain Models', 'Package/Advisory\nStatistics/Version', 3.5),
        ('Detector', 'Cross-OS\nDetection', 7.5),
        ('Installer', 'Auto-Install\nPHP+Composer', 11.5),
        ('Utilities', 'FS/HTTP\nMock Helpers', 15.5),
        ('Composerutils', 'Shared Utils', 19.5),
    ]
    for t, s, fx in f_items:
        draw_card(ax, fx, 3.2, 3.2, 0.75, t, s,
                  color=PAL['text2'], bg=PAL['surface2'], title_fontsize=8.5, sub_fontsize=6.5,
                  shadow=False, border_width=1, radius=0.12)

    # Stats box
    draw_card(ax, 4, 0.8, 7, 1.2, 'Project Stats',
              '20 API | 234 CLI | 20 Categories\n50+ Commands | 450+ Tests | 12 Guides',
              color=PAL['primary'], bg=PAL['primary_bg'], title_fontsize=11, sub_fontsize=8.5, shadow=True)

    plt.tight_layout(pad=0.3)
    plt.savefig('/home/cc11001100/github/scagogogo/composer-skills/docs/images/feature-mindmap.png',
                dpi=150, bbox_inches='tight', facecolor=PAL['bg'])
    print("Feature map saved!")


# ══════════════════════════════════════════════════════════════
# 4. SDK Comparison
# ══════════════════════════════════════════════════════════════
def generate_sdk_comparison():
    fig, ax = plt.subplots(1, 1, figsize=(16, 7))
    ax.set_xlim(0, 16)
    ax.set_ylim(0, 7)
    ax.set_aspect('equal')
    ax.axis('off')
    fig.patch.set_facecolor(PAL['bg'])
    draw_gradient_bg(ax, (0, 16), (0, 7), '#EEF2FF', '#F5F3FF')

    # Left card
    lc = FancyBboxPatch(
        (0.5, 0.5), 6.5, 6,
        boxstyle="round,pad=0.04,rounding_size=0.3",
        facecolor=PAL['surface'], edgecolor=PAL['info'],
        linewidth=2.5, alpha=0.95, zorder=2
    )
    ax.add_patch(lc)
    lh = FancyBboxPatch(
        (0.5, 5.4), 6.5, 1.1,
        boxstyle="round,pad=0.02,rounding_size=0.3",
        facecolor=PAL['info'], edgecolor='none', alpha=0.95, zorder=3
    )
    ax.add_patch(lh)
    ax.text(3.75, 6.2, 'Packagist API SDK', ha='center', va='center',
            fontsize=16, color='white', weight='bold', zorder=4)
    ax.text(3.75, 5.8, 'pkg/client | pkg/repository', ha='center', va='center',
            fontsize=9, color='#BFDBFE', zorder=4)

    left_items = [
        'Search packages & browse registry',
        'Get package details & statistics',
        'Security advisories & CVE tracking',
        'List packages by vendor / type',
        'Pure Go -- no PHP required',
    ]
    for i, item in enumerate(left_items):
        c = PAL['success'] if i < 4 else PAL['info']
        ax.plot(1.3, 4.8 - i * 0.65, 'o', color=c, markersize=5, zorder=4)
        ax.text(3.75, 4.8 - i * 0.65, item, ha='center', va='center',
                fontsize=10, color=PAL['text'], zorder=4)

    # Right card
    rc = FancyBboxPatch(
        (9, 0.5), 6.5, 6,
        boxstyle="round,pad=0.04,rounding_size=0.3",
        facecolor=PAL['surface'], edgecolor=PAL['secondary'],
        linewidth=2.5, alpha=0.95, zorder=2
    )
    ax.add_patch(rc)
    rh = FancyBboxPatch(
        (9, 5.4), 6.5, 1.1,
        boxstyle="round,pad=0.02,rounding_size=0.3",
        facecolor=PAL['secondary'], edgecolor='none', alpha=0.95, zorder=3
    )
    ax.add_patch(rh)
    ax.text(12.25, 6.2, 'Composer CLI SDK', ha='center', va='center',
            fontsize=16, color='white', weight='bold', zorder=4)
    ax.text(12.25, 5.8, 'pkg/composer', ha='center', va='center',
            fontsize=9, color='#DDD6FE', zorder=4)

    right_items = [
        '234 methods across 20 categories',
        'Install / update / require / remove',
        'Security audit & vulnerability scan',
        'Project management & scripts',
        'Requires PHP 7.4+ & Composer 2.0+',
    ]
    for i, item in enumerate(right_items):
        c = PAL['success'] if i < 4 else PAL['secondary']
        ax.plot(9.8, 4.8 - i * 0.65, 'o', color=c, markersize=5, zorder=4)
        ax.text(12.25, 4.8 - i * 0.65, item, ha='center', va='center',
                fontsize=10, color=PAL['text'], zorder=4)

    # VS badge
    vc = plt.Circle((8, 3.5), 0.55, color=PAL['text'], alpha=0.9, zorder=5)
    ax.add_patch(vc)
    ax.text(8, 3.5, 'VS', ha='center', va='center', fontsize=13,
            color='white', weight='bold', zorder=6)

    plt.tight_layout(pad=0.5)
    plt.savefig('/home/cc11001100/github/scagogogo/composer-skills/docs/images/sdk-comparison.png',
                dpi=150, bbox_inches='tight', facecolor=PAL['bg'])
    print("SDK comparison saved!")


# ══════════════════════════════════════════════════════════════
# 5. Auto-Install Flow
# ══════════════════════════════════════════════════════════════
def generate_auto_install_flow():
    fig, ax = plt.subplots(1, 1, figsize=(16, 9))
    ax.set_xlim(0, 16)
    ax.set_ylim(0, 9)
    ax.set_aspect('equal')
    ax.axis('off')
    fig.patch.set_facecolor(PAL['bg'])
    draw_gradient_bg(ax, (0, 16), (0, 9), '#ECFDF5', '#EEF2FF')

    ax.text(8, 8.5, 'Auto-Install Flow', ha='center', va='center',
            fontsize=22, color=PAL['text'], weight='bold')
    ax.text(8, 8.05, 'Detect > Check > Install > Verify > Ready',
            ha='center', va='center', fontsize=10, color=PAL['text3'])

    steps = [
        (2, 6.5, '1', 'Detect', 'Is Composer\ninstalled?', PAL['info'], PAL['info_bg']),
        (6, 6.5, '2', 'Check PHP', 'Is PHP\navailable?', PAL['warning'], PAL['warning_bg']),
        (10, 6.5, '3', 'Install', 'Install missing\ndependencies', PAL['accent'], PAL['accent_bg']),
        (14, 6.5, '4', 'Verify', 'Detect & validate\ninstallation', PAL['success'], PAL['success_bg']),
    ]
    for x, y, num, title, desc, color, bg in steps:
        # Step number circle
        nc = plt.Circle((x - 0.9, y + 0.55), 0.25, color=color, alpha=0.85, zorder=5)
        ax.add_patch(nc)
        ax.text(x - 0.9, y + 0.55, num, ha='center', va='center',
                fontsize=11, color='white', weight='bold', zorder=6)
        draw_card(ax, x, y, 2.8, 1.5, title, desc,
                  color=color, bg=bg, title_fontsize=11, sub_fontsize=8, shadow=True)

    # Main flow arrows
    draw_connector(ax, 3.4, 6.5, 4.6, 6.5, color=PAL['info'], lw=2)
    draw_connector(ax, 7.4, 6.5, 8.6, 6.5, color=PAL['warning'], lw=2)
    draw_connector(ax, 11.4, 6.5, 12.6, 6.5, color=PAL['accent'], lw=2)

    # Shortcuts
    draw_card(ax, 2, 4.1, 2.4, 0.7, 'Already installed!', 'Skip to step 4',
              color=PAL['success'], bg=PAL['success_bg'], title_fontsize=9, sub_fontsize=7, shadow=False)
    draw_connector(ax, 2, 5.75, 2, 4.45, color=PAL['success'], lw=1.5, style='--')
    draw_connector(ax, 3.2, 4.1, 12.6, 6.1, color=PAL['success'], lw=1.5, style='--')

    draw_card(ax, 6, 4.1, 2.4, 0.7, 'PHP available', 'Skip to step 3',
              color=PAL['warning'], bg=PAL['warning_bg'], title_fontsize=9, sub_fontsize=7, shadow=False)
    draw_connector(ax, 6, 5.75, 6, 4.45, color=PAL['warning'], lw=1.5, style='--')
    draw_connector(ax, 7.2, 4.1, 8.6, 6.1, color=PAL['warning'], lw=1.5, style='--')

    # Result
    draw_card(ax, 8, 2.0, 5, 1.1, 'Ready!',
              'Composer instance created\nOne import: github.com/scagogogo/composer-skills',
              color=PAL['primary'], bg=PAL['primary_bg'], title_fontsize=13, sub_fontsize=9, shadow=True)
    draw_connector(ax, 14, 5.75, 10.5, 2.55, color=PAL['primary'], lw=2, curved=True)

    plt.tight_layout(pad=0.5)
    plt.savefig('/home/cc11001100/github/scagogogo/composer-skills/docs/images/auto-install-flow.png',
                dpi=150, bbox_inches='tight', facecolor=PAL['bg'])
    print("Auto-install flow saved!")


# ══════════════════════════════════════════════════════════════
# 6. Platform Matrix
# ══════════════════════════════════════════════════════════════
def generate_platform_matrix():
    fig, ax = plt.subplots(1, 1, figsize=(13, 6))
    ax.set_xlim(0, 13)
    ax.set_ylim(0, 6)
    ax.set_aspect('equal')
    ax.axis('off')
    fig.patch.set_facecolor(PAL['bg'])

    ax.text(6.5, 5.6, 'Platform Support', ha='center', va='center',
            fontsize=20, color=PAL['text'], weight='bold')

    headers = ['Feature', 'Linux', 'macOS', 'Windows']
    col_x = [2, 5, 8, 11]
    h_colors = [PAL['primary'], PAL['accent'], PAL['info'], PAL['success']]
    h_bgs = [PAL['primary_bg'], PAL['accent_bg'], PAL['info_bg'], PAL['success_bg']]

    for i, (h, x, c, b) in enumerate(zip(headers, col_x, h_colors, h_bgs)):
        pill = FancyBboxPatch(
            (x - 1.2, 4.85), 2.4, 0.5,
            boxstyle="round,pad=0.02,rounding_size=0.12",
            facecolor=b, edgecolor=c, linewidth=2, alpha=0.95, zorder=3
        )
        ax.add_patch(pill)
        ax.text(x, 5.1, h, ha='center', va='center',
                fontsize=11, color=c, weight='bold', zorder=4)

    rows = [
        ('Composer Detection', 'Yes', 'Yes', 'Yes'),
        ('Auto-Install', 'Yes', 'Yes', 'Yes'),
        ('PHP Auto-Install', 'Yes', 'Yes', '--'),
        ('Package Manager', 'apt/dnf/pacman\napk/zypper', 'Homebrew', '--'),
        ('Direct Download', 'Yes', 'Yes', 'Yes'),
    ]
    for ri, (feat, linux, macos, windows) in enumerate(rows):
        y = 4.2 - ri * 0.85
        rbg = FancyBboxPatch(
            (0.5, y - 0.3), 12, 0.6,
            boxstyle="round,pad=0.02,rounding_size=0.1",
            facecolor=PAL['surface'] if ri % 2 == 0 else PAL['surface2'],
            edgecolor='none', alpha=0.7, zorder=1
        )
        ax.add_patch(rbg)

        ax.text(col_x[0], y, feat, ha='center', va='center',
                fontsize=9.5, color=PAL['text'], weight='bold', zorder=4)

        for val, x in zip([linux, macos, windows], col_x[1:]):
            if val == 'Yes':
                ax.text(x, y, 'Yes', ha='center', va='center',
                        fontsize=11, color=PAL['success'], weight='bold', zorder=4)
            elif val == '--':
                ax.text(x, y, '--', ha='center', va='center',
                        fontsize=11, color=PAL['text3'], zorder=4)
            else:
                ax.text(x, y, val, ha='center', va='center',
                        fontsize=7, color=PAL['text2'], zorder=4, linespacing=1.2)

    ax.text(6.5, 0.35, 'Yes = Supported    -- = Not applicable',
            ha='center', va='center', fontsize=9, color=PAL['text3'])

    plt.tight_layout(pad=0.5)
    plt.savefig('/home/cc11001100/github/scagogogo/composer-skills/docs/images/platform-matrix.png',
                dpi=150, bbox_inches='tight', facecolor=PAL['bg'])
    print("Platform matrix saved!")


# ══════════════════════════════════════════════════════════════
# 7. Security Features
# ══════════════════════════════════════════════════════════════
def generate_security_diagram():
    fig, ax = plt.subplots(1, 1, figsize=(14, 8))
    ax.set_xlim(0, 14)
    ax.set_ylim(0, 8)
    ax.set_aspect('equal')
    ax.axis('off')
    fig.patch.set_facecolor(PAL['bg'])
    draw_gradient_bg(ax, (0, 14), (0, 8), '#FFF1F2', '#F5F3FF')

    ax.text(7, 7.6, 'Security-First Design', ha='center', va='center',
            fontsize=22, color=PAL['text'], weight='bold')

    # Center hub
    hub = FancyBboxPatch(
        (5.3, 4.65), 3.4, 1.5,
        boxstyle="round,pad=0.04,rounding_size=0.25",
        facecolor=PAL['accent'], edgecolor=PAL['accent'],
        linewidth=2, alpha=0.95, zorder=3
    )
    ax.add_patch(hub)
    ax.text(7, 5.55, 'Security', ha='center', va='center',
            fontsize=16, color='white', weight='bold', zorder=4)
    ax.text(7, 5.1, 'Audit', ha='center', va='center',
            fontsize=11, color='#FFE4E6', zorder=4)

    spokes = [
        (2, 6.3, 'Vulnerability\nScan', 'HasVulnerabilities\nGetHighSeverityVulns', PAL['accent'], PAL['accent_bg']),
        (12, 6.3, 'Packagist\nAdvisories', 'GetSecurityAdvisories\nGetAdvisoriesForPkg', PAL['warning'], PAL['warning_bg']),
        (2, 3.0, 'Schema\nValidation', 'Validate\nValidateStrict\nValidateSchema', PAL['success'], PAL['success_bg']),
        (12, 3.0, 'Platform\nRequirements', 'CheckPlatform\nGetPHPVersion\nHasExtension', PAL['info'], PAL['info_bg']),
        (7, 1.2, 'License\nCompliance', 'Licenses\nCheckLicenses\nLicensesWithFormat', PAL['secondary'], PAL['secondary_bg']),
    ]
    for x, y, t, s, c, b in spokes:
        draw_card(ax, x, y, 3.0, 1.5, t, s,
                  color=c, bg=b, title_fontsize=10, sub_fontsize=7,
                  shadow=True, border_width=1.5)
        draw_connector(ax, 7, 4.65, x, y + 0.75, color=c, lw=1.5, arrow=False)

    plt.tight_layout(pad=0.5)
    plt.savefig('/home/cc11001100/github/scagogogo/composer-skills/docs/images/security-features.png',
                dpi=150, bbox_inches='tight', facecolor=PAL['bg'])
    print("Security features saved!")


if __name__ == '__main__':
    generate_header_badge()
    generate_architecture_diagram()
    generate_feature_mindmap()
    generate_sdk_comparison()
    generate_auto_install_flow()
    generate_platform_matrix()
    generate_security_diagram()
    print("\nAll images generated!")
