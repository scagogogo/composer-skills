import React from 'react'
import { useTranslation } from 'react-i18next'
import { Tabs } from 'antd'
import { DownloadOutlined, CloudOutlined, CodeOutlined, ThunderboltOutlined } from '@ant-design/icons'
import SectionTitle from '../components/SectionTitle'
import CodeBlock from '../components/CodeBlock'
import { RevealSection } from '../components/ScrollReveal'

const QuickStartSection: React.FC = () => {
  const { t } = useTranslation()

  const tabItems = [
    {
      key: 'packagist',
      label: (
        <span style={{ display: 'flex', alignItems: 'center', gap: 6 }}>
          <CloudOutlined /> {t('quickStart.tabPackagist')}
        </span>
      ),
      children: <CodeBlock code={t('quickStart.packagistCode')} />,
    },
    {
      key: 'composer',
      label: (
        <span style={{ display: 'flex', alignItems: 'center', gap: 6 }}>
          <CodeOutlined /> {t('quickStart.tabComposer')}
        </span>
      ),
      children: <CodeBlock code={t('quickStart.composerCode')} />,
    },
    {
      key: 'autoInstall',
      label: (
        <span style={{ display: 'flex', alignItems: 'center', gap: 6 }}>
          <ThunderboltOutlined /> {t('quickStart.tabAutoInstall')}
        </span>
      ),
      children: <CodeBlock code={t('quickStart.autoInstallCode')} />,
    },
  ]

  return (
    <section id="quickStart">
      <div style={{ maxWidth: 1100, margin: '0 auto' }}>
        <SectionTitle title={t('quickStart.title')} subtitle={t('quickStart.subtitle')} />

        {/* Install command — prominent box */}
        <RevealSection>
          <div style={{ maxWidth: 800, margin: '0 auto 36px' }}>
            <div
              style={{
                display: 'flex',
                alignItems: 'center',
                gap: 10,
                marginBottom: 10,
              }}
            >
              <DownloadOutlined style={{ color: '#2563EB', fontSize: 16 }} />
              <span style={{ fontWeight: 700, fontSize: 15, color: '#0F172A' }}>
                {t('quickStart.install')}
              </span>
            </div>
            <CodeBlock code={t('quickStart.installCode')} language="bash" />
          </div>
        </RevealSection>

        {/* Code tabs */}
        <RevealSection delay={0.1}>
          <div style={{ maxWidth: 900, margin: '0 auto 40px' }}>
            <Tabs
              defaultActiveKey="packagist"
              items={tabItems}
              centered
            />
          </div>
        </RevealSection>

        {/* Convenience methods */}
        <RevealSection delay={0.15}>
          <div style={{ maxWidth: 800, margin: '0 auto' }}>
            <div
              style={{
                border: '1px solid #E2E8F0',
                borderRadius: 4,
                overflow: 'hidden',
              }}
            >
              <div
                style={{
                  padding: '10px 16px',
                  background: '#F8FAFC',
                  borderBottom: '1px solid #E2E8F0',
                  fontWeight: 700,
                  fontSize: 14,
                  color: '#0F172A',
                }}
              >
                {t('quickStart.convenienceTitle')}
              </div>
              <div style={{ padding: '0 12px 12px' }}>
                <CodeBlock code={t('quickStart.convenienceCode')} />
              </div>
            </div>
          </div>
        </RevealSection>
      </div>
    </section>
  )
}

export default QuickStartSection
