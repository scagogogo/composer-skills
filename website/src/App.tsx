import React from 'react'
import { ConfigProvider, App as AntApp } from 'antd'
import { themeConfig } from './theme/token'
import './i18n'
import Layout from './layouts/Layout'
import HeroSection from './sections/HeroSection'
import ProblemSection from './sections/ProblemSection'
import FeatureSection from './sections/FeatureSection'
import ArchitectureSection from './sections/ArchitectureSection'
import SDKComparisonSection from './sections/SDKComparisonSection'
import SecuritySection from './sections/SecuritySection'
import AutoInstallSection from './sections/AutoInstallSection'
import CoverageSection from './sections/CoverageSection'
import QuickStartSection from './sections/QuickStartSection'
import UseCasesSection from './sections/UseCasesSection'
import TutorialsSection from './sections/TutorialsSection'
import ShowcaseSection from './sections/ShowcaseSection'

const App: React.FC = () => {
  return (
    <ConfigProvider theme={themeConfig}>
      <AntApp>
        <Layout>
          <HeroSection />
          <ProblemSection />
          <FeatureSection />
          <ArchitectureSection />
          <SDKComparisonSection />
          <SecuritySection />
          <AutoInstallSection />
          <QuickStartSection />
          <TutorialsSection />
          <ShowcaseSection />
          <CoverageSection />
          <UseCasesSection />
        </Layout>
      </AntApp>
    </ConfigProvider>
  )
}

export default App
