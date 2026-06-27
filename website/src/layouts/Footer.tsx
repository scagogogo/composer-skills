import React from 'react'
import { useTranslation } from 'react-i18next'
import { Layout as AntLayout, Typography, Space } from 'antd'

const { Footer: AntFooter } = AntLayout
const { Text, Link } = Typography

const Footer: React.FC = () => {
  const { t } = useTranslation()

  return (
    <AntFooter
      style={{
        background: 'linear-gradient(180deg, #0F172A, #1E293B)',
        color: '#CBD5E1',
        padding: '56px 24px 28px',
      }}
    >
      <div style={{ maxWidth: 1100, margin: '0 auto' }}>
        <div
          style={{
            display: 'grid',
            gridTemplateColumns: 'repeat(auto-fit, minmax(200px, 1fr))',
            gap: 40,
            marginBottom: 40,
          }}
        >
          <div>
            <div style={{ display: 'flex', alignItems: 'center', gap: 10, marginBottom: 16 }}>
              <div
                style={{
                  width: 32,
                  height: 32,
                  borderRadius: 8,
                  background: 'linear-gradient(135deg, #6366F1, #818CF8)',
                  display: 'flex',
                  alignItems: 'center',
                  justifyContent: 'center',
                  color: '#fff',
                  fontWeight: 800,
                  fontSize: 14,
                  fontFamily: 'monospace',
                }}
              >
                CS
              </div>
              <Text strong style={{ color: '#F8FAFC', fontSize: 17 }}>Composer Skills</Text>
            </div>
            <Text style={{ color: '#94A3B8', lineHeight: 1.6 }}>{t('footer.description')}</Text>
          </div>
          <div>
            <Text strong style={{ color: '#F8FAFC', fontSize: 15, display: 'block', marginBottom: 16 }}>
              {t('footer.resources')}
            </Text>
            <Space direction="vertical" size={10}>
              <Link
                href="https://github.com/scagogogo/composer-skills/blob/main/docs/skills/01-getting-started.md"
                style={{ color: '#94A3B8', fontSize: 14 }}
              >
                {t('footer.docGettingStarted')}
              </Link>
              <Link
                href="https://github.com/scagogogo/composer-skills/blob/main/docs/skills/02-packagist-api.md"
                style={{ color: '#94A3B8', fontSize: 14 }}
              >
                {t('footer.docPackagist')}
              </Link>
              <Link
                href="https://github.com/scagogogo/composer-skills/blob/main/docs/skills/05-security.md"
                style={{ color: '#94A3B8', fontSize: 14 }}
              >
                {t('footer.docSecurity')}
              </Link>
              <Link
                href="https://github.com/scagogogo/composer-skills/blob/main/docs/skills/11-cli-reference.md"
                style={{ color: '#94A3B8', fontSize: 14 }}
              >
                {t('footer.docCLI')}
              </Link>
            </Space>
          </div>
          <div>
            <Text strong style={{ color: '#F8FAFC', fontSize: 15, display: 'block', marginBottom: 16 }}>
              {t('footer.community')}
            </Text>
            <Space direction="vertical" size={10}>
              <Link href="https://github.com/scagogogo/composer-skills" style={{ color: '#94A3B8', fontSize: 14 }}>
                {t('footer.github')}
              </Link>
              <Link href="https://pkg.go.dev/github.com/scagogogo/composer-skills" style={{ color: '#94A3B8', fontSize: 14 }}>
                {t('footer.goReference')}
              </Link>
              <Link href="https://goreportcard.com/report/github.com/scagogogo/composer-skills" style={{ color: '#94A3B8', fontSize: 14 }}>
                {t('footer.goReport')}
              </Link>
            </Space>
          </div>
          <div>
            <Text strong style={{ color: '#F8FAFC', fontSize: 15, display: 'block', marginBottom: 16 }}>
              {t('footer.acknowledgments')}
            </Text>
            <Space direction="vertical" size={10}>
              <Link href="https://packagist.org" style={{ color: '#94A3B8', fontSize: 14 }}>
                {t('footer.packagist')}
              </Link>
              <Link href="https://getcomposer.org" style={{ color: '#94A3B8', fontSize: 14 }}>
                {t('footer.composer')}
              </Link>
            </Space>
          </div>
        </div>
        <div
          style={{
            borderTop: '1px solid rgba(255, 255, 255, 0.08)',
            paddingTop: 20,
            textAlign: 'center',
            display: 'flex',
            justifyContent: 'space-between',
            alignItems: 'center',
            flexWrap: 'wrap',
            gap: 12,
          }}
        >
          <Text style={{ color: '#64748B', fontSize: 13 }}>
            {t('footer.copyright')}
          </Text>
          <Text style={{ color: '#475569', fontSize: 12 }}>
            Built with ❤️ using Go + React + Ant Design
          </Text>
        </div>
      </div>
    </AntFooter>
  )
}

export default Footer
