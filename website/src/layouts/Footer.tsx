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
        background: '#0F172A',
        color: '#94A3B8',
        padding: '48px 24px 24px',
      }}
    >
      <div style={{ maxWidth: 1100, margin: '0 auto' }}>
        <div
          style={{
            display: 'grid',
            gridTemplateColumns: 'repeat(auto-fit, minmax(180px, 1fr))',
            gap: 32,
            marginBottom: 32,
          }}
        >
          <div>
            <div style={{ display: 'flex', alignItems: 'center', gap: 8, marginBottom: 12 }}>
              <img
                src={`${import.meta.env.BASE_URL}logo.svg`}
                alt="Composer Skills"
                style={{ width: 24, height: 24 }}
              />
              <Text strong style={{ color: '#F8FAFC', fontSize: 15 }}>Composer Skills</Text>
            </div>
            <Text style={{ color: '#64748B', lineHeight: 1.6, fontSize: 14 }}>{t('footer.description')}</Text>
          </div>
          <div>
            <Text strong style={{ color: '#F8FAFC', fontSize: 14, display: 'block', marginBottom: 12 }}>
              {t('footer.resources')}
            </Text>
            <Space direction="vertical" size={8}>
              <Link
                href="https://github.com/scagogogo/composer-skills/blob/main/docs/skills/01-getting-started.md"
                style={{ color: '#64748B', fontSize: 13 }}
              >
                {t('footer.docGettingStarted')}
              </Link>
              <Link
                href="https://github.com/scagogogo/composer-skills/blob/main/docs/skills/02-packagist-api.md"
                style={{ color: '#64748B', fontSize: 13 }}
              >
                {t('footer.docPackagist')}
              </Link>
              <Link
                href="https://github.com/scagogogo/composer-skills/blob/main/docs/skills/05-security.md"
                style={{ color: '#64748B', fontSize: 13 }}
              >
                {t('footer.docSecurity')}
              </Link>
              <Link
                href="https://github.com/scagogogo/composer-skills/blob/main/docs/skills/11-cli-reference.md"
                style={{ color: '#64748B', fontSize: 13 }}
              >
                {t('footer.docCLI')}
              </Link>
            </Space>
          </div>
          <div>
            <Text strong style={{ color: '#F8FAFC', fontSize: 14, display: 'block', marginBottom: 12 }}>
              {t('footer.community')}
            </Text>
            <Space direction="vertical" size={8}>
              <Link href="https://github.com/scagogogo/composer-skills" style={{ color: '#64748B', fontSize: 13 }}>
                {t('footer.github')}
              </Link>
              <Link href="https://pkg.go.dev/github.com/scagogogo/composer-skills" style={{ color: '#64748B', fontSize: 13 }}>
                {t('footer.goReference')}
              </Link>
              <Link href="https://goreportcard.com/report/github.com/scagogogo/composer-skills" style={{ color: '#64748B', fontSize: 13 }}>
                {t('footer.goReport')}
              </Link>
            </Space>
          </div>
          <div>
            <Text strong style={{ color: '#F8FAFC', fontSize: 14, display: 'block', marginBottom: 12 }}>
              {t('footer.acknowledgments')}
            </Text>
            <Space direction="vertical" size={8}>
              <Link href="https://packagist.org" style={{ color: '#64748B', fontSize: 13 }}>
                {t('footer.packagist')}
              </Link>
              <Link href="https://getcomposer.org" style={{ color: '#64748B', fontSize: 13 }}>
                {t('footer.composer')}
              </Link>
            </Space>
          </div>
        </div>
        <div
          style={{
            borderTop: '1px solid #1E293B',
            paddingTop: 16,
            textAlign: 'center',
          }}
        >
          <Text style={{ color: '#475569', fontSize: 12 }}>
            {t('footer.copyright')}
          </Text>
        </div>
      </div>
    </AntFooter>
  )
}

export default Footer
