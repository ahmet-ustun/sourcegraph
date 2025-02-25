package com.sourcegraph.cody.config

import com.intellij.openapi.progress.ProgressIndicator
import com.intellij.openapi.ui.ValidationInfo
import com.intellij.ui.layout.LayoutBuilder
import com.intellij.ui.layout.panel
import com.intellij.util.ui.JBEmptyBorder
import com.intellij.util.ui.UIUtil
import com.sourcegraph.cody.api.SourcegraphApiRequestExecutor
import javax.swing.JComponent
import javax.swing.JPanel

internal abstract class CodyCredentialsUi {
    abstract fun getPreferredFocusableComponent(): JComponent?
    abstract fun getValidator(): Validator
    abstract fun createExecutor(): SourcegraphApiRequestExecutor
    abstract fun acquireLoginAndToken(
        server: SourcegraphServerPath,
        executor: SourcegraphApiRequestExecutor,
        indicator: ProgressIndicator
    ): Pair<String,String>

    abstract fun handleAcquireError(error: Throwable): ValidationInfo
    abstract fun setBusy(busy: Boolean)

    var footer: LayoutBuilder.() -> Unit = { }

    fun getPanel(): JPanel =
        panel {
            centerPanel()
            footer()
        }.apply {
            // Border is required to have more space - otherwise there could be issues with focus ring.
            // `getRegularPanelInsets()` is used to simplify border calculation for dialogs where this panel is used.
            border = JBEmptyBorder(UIUtil.getRegularPanelInsets())
        }

    protected abstract fun LayoutBuilder.centerPanel()
}
